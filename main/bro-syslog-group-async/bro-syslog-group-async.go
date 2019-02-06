package main

import (
	"flag"
	"fmt"
	input "github.com/xemoe/go-syslog-report/input"
	validators "github.com/xemoe/go-syslog-report/validators"
	workers "github.com/xemoe/go-syslog-report/workers"
	"log"
	"strconv"
	"strings"
)

var (
	filename       = flag.String("f", "syslog.log.gz", "the file to process")
	group          = flag.String("g", "id.orig_h,facility", "the fields to process")
	reduceParallel = flag.String("t", "4", "number of reduce workers")
	mapperParallel = flag.String("w", "4", "number of mapper workers")
)

func main() {

	flag.Parse()
	log.Printf("Read syslog from input:%s", *filename)

	files := strings.Split(*filename, ",")
	_, err := validators.ValidateFilesIsExists(files)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read syslog from files:%q", files)

	reduceWorkers, err := strconv.Atoi(*reduceParallel)
	if err != nil {
		log.Fatal(err)
	}
	mapperWorkers, err := strconv.Atoi(*mapperParallel)
	if err != nil {
		log.Fatal(err)
	}

	skipIndex := input.ParseIndexFields(*group)
	result := workers.GroupCountMultiplesAsync(reduceWorkers, mapperWorkers, files, skipIndex)

	for _, v := range result {
		fmt.Printf("%s|%d\n", v.Key, v.Value)
	}
}
