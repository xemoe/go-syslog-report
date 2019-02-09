package workers

import (
	"bufio"
	gzip "github.com/klauspost/pgzip"
	cache "github.com/xemoe/go-syslog-report/cache"
	mapper "github.com/xemoe/go-syslog-report/mapper"
	types "github.com/xemoe/go-syslog-report/types"
	"log"
	"os"
	"reflect"
	"sort"
	"sync"
)

var CacheDir = "/tmp"
var CachePrefix = "workers"

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

	cache.CacheDir = CacheDir
	cache.CachePrefix = CachePrefix

	indexhash, filenamehash, filechecksum, result, err := cache.GetCache(filename, index)
	if err == nil {
		return result
	}

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
	result = make(map[string]int)
	ref := reflect.ValueOf(index)
	numfields := ref.NumField()

	for scanner.Scan() {
		line := scanner.Text()
		uniqkey, isok := mapper.SyslogGroupMapper(ref, numfields, line)
		if isok {
			result[uniqkey] += 1
		}
	}

	err = cache.SaveCache(indexhash, filenamehash, filechecksum, result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
