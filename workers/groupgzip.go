package workers

import (
	"bufio"
	gzip "github.com/klauspost/pgzip"
	mapper "github.com/xemoe/go-syslog-report/mapper"
	types "github.com/xemoe/go-syslog-report/types"
	"log"
	"os"
	"reflect"
	"sort"
	"sync"
)

func GroupCountMultiplesSync(files []string, index types.SyslogLineIndex) []types.GroupCountSlices {

	result := map[string]int{}

	for i := 0; i < len(files); i++ {
		one := GroupCount(files[i], index)
		for uniqkey, uniqval := range one {
			result[uniqkey] += uniqval
		}
	}

	myslices := []types.GroupCountSlices{}
	for uniqkey, uniqval := range result {
		myslices = append(myslices, types.GroupCountSlices{uniqkey, uniqval})
	}

	//
	// Sort
	//
	sort.Slice(myslices, func(i, j int) bool { return myslices[i].Value > myslices[j].Value })

	return myslices
}

func GroupCountMultiples(files []string, index types.SyslogLineIndex) []types.GroupCountSlices {

	wg := new(sync.WaitGroup)
	result := map[string]int{}
	crs := make(chan map[string]int, len(files))

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(filename string, index types.SyslogLineIndex, crs chan<- map[string]int) {
			crs <- GroupCount(filename, index)
			defer wg.Done()
		}(files[i], index, crs)
	}

	go func() {
		wg.Wait()
		close(crs)
	}()

	//
	// Reduces
	//
	for one := range crs {
		for uniqkey, uniqval := range one {
			result[uniqkey] += uniqval
		}
	}

	myslices := []types.GroupCountSlices{}
	for uniqkey, uniqval := range result {
		myslices = append(myslices, types.GroupCountSlices{uniqkey, uniqval})
	}

	//
	// Sort
	//
	sort.Slice(myslices, func(i, j int) bool { return myslices[i].Value > myslices[j].Value })

	return myslices
}

func GroupCount(filename string, index types.SyslogLineIndex) map[string]int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()

	scanner := bufio.NewScanner(gz)

	//
	// Processing result
	//
	result := make(map[string]int)
	ref := reflect.ValueOf(index)
	numfields := ref.NumField()

	for scanner.Scan() {
		line := scanner.Text()
		uniqkey, isok := mapper.SyslogGroupMapper(ref, numfields, line)
		if isok {
			result[uniqkey] += 1
		}
	}

	return result
}

func GroupCountWithChan(filename string, index types.SyslogLineIndex) map[string]int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()

	scanner := bufio.NewScanner(gz)

	//
	// Processing result
	//
	ref := reflect.ValueOf(index)
	numfields := ref.NumField()

	//
	// Create jobs
	//
	jobs := make(chan string)
	results := make(chan string)
	wg := new(sync.WaitGroup)

	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go mapping(ref, numfields, jobs, results, wg)
	}

	for scanner.Scan() {
		jobs <- scanner.Text()
	}

	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	return reduce(results)
}

func mapping(ref reflect.Value, numfields int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range jobs {
		item, isok := mapper.SyslogGroupMapper(ref, numfields, line)
		if isok {
			results <- item
		}
	}
}

func reduce(results <-chan string) map[string]int {
	sum := map[string]int{}
	for v := range results {
		sum[v] += 1
	}

	return sum
}
