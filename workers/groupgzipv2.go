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

type Map struct {
	sync.Mutex
	m map[string]int
}

func GroupCountMultiplesAsync(reduceParallel int, mapperParallel int, files []string, index types.SyslogLineIndex) []types.GroupCountSlices {

	//
	// Data index
	//
	ref := reflect.ValueOf(index)
	numfields := ref.NumField()

	var sum = &Map{
		m: map[string]int{},
	}

	reduceChan, reduceWaitGroup := startReduce(reduceParallel, sum)
	mapperChan, mapperWaitGroup := startMapper(mapperParallel, reduceChan, ref, numfields)
	readerWaitGroup := startReaders(mapperChan, files)

	readerWaitGroup.Wait()
	close(mapperChan)

	mapperWaitGroup.Wait()
	close(reduceChan)

	reduceWaitGroup.Wait()

	myslices := []types.GroupCountSlices{}
	for uniqkey, uniqval := range (*sum).m {
		myslices = append(myslices, types.GroupCountSlices{uniqkey, uniqval})
	}

	//
	// Sort
	//
	sort.Slice(myslices, func(i, j int) bool { return myslices[i].Value > myslices[j].Value })

	return myslices
}

//
// Parallel reducer
//
func startReduce(reduceParallel int, sum *Map) (chan string, *sync.WaitGroup) {
	reduceWaitGroup := &sync.WaitGroup{}
	reduceChan := make(chan string)

	for i := 0; i < reduceParallel; i++ {
		reduceWaitGroup.Add(1)
		go reducer(sum, reduceChan, reduceWaitGroup)
	}

	return reduceChan, reduceWaitGroup
}

//
// Output reducer
//
func reducer(sum *Map, reduceChan chan string, reduceWaitGroup *sync.WaitGroup) {
	for v := range reduceChan {
		(*sum).Lock()
		(*sum).m[v] += 1
		(*sum).Unlock()
	}

	reduceWaitGroup.Done()
}

//
// Parallel mapper
//
func startMapper(mapperParallel int, reduceChan chan string, ref reflect.Value, numfields int) (chan string, *sync.WaitGroup) {
	mapperWaitGroup := &sync.WaitGroup{}
	mapperChan := make(chan string)

	for i := 0; i < mapperParallel; i++ {
		mapperWaitGroup.Add(1)
		go mapperjob(reduceChan, mapperChan, mapperWaitGroup, ref, numfields)
	}

	return mapperChan, mapperWaitGroup
}

func mapperjob(reduceChan chan string, mapperChan chan string, mapperWaitGroup *sync.WaitGroup, ref reflect.Value, numfields int) {
	for line := range mapperChan {
		item, ok := mapper.SyslogGroupMapper(ref, numfields, line)
		if ok {
			reduceChan <- item
		}

	}
	mapperWaitGroup.Done()
}

//
// Parallel read chan
//
func startReaders(mapperChan chan string, files []string) *sync.WaitGroup {
	readerWaitGroup := &sync.WaitGroup{}
	for i := 0; i < len(files); i++ {
		readerWaitGroup.Add(1)
		go reader(mapperChan, readerWaitGroup, files[i])
	}

	return readerWaitGroup
}

func reader(mapperChan chan string, readerWaitGroup *sync.WaitGroup, filename string) {

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

	for scanner.Scan() {
		mapperChan <- scanner.Text()
	}

	readerWaitGroup.Done()
}
