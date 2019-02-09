package main

import (
	"flag"
	"fmt"
	input "github.com/xemoe/go-syslog-report/input"
	validators "github.com/xemoe/go-syslog-report/validators"
	workers "github.com/xemoe/go-syslog-report/workers"
	"log"
	"strings"
)

var (
	filename = flag.String("f", "syslog.log.gz", "the file to process")
	group    = flag.String("g", "id.orig_h,facility", "the fields to process")
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

	workers.CacheDir = "/tmp"
	workers.CachePrefix = "bdtest"

	skipIndex := input.ParseIndexFields(*group)
	result := workers.GroupCountMultiples(files, skipIndex)

	for _, v := range result {
		fmt.Printf("%s|%d\n", v.Key, v.Value)
	}
}
