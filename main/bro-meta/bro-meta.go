package main

import (
	"flag"
	workers "github.com/xemoe/go-syslog-report/workers"
	"log"
)

var (
	filename = flag.String("f", "syslog.log.gz", "the file to process")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Read meta info from file:%s", *filename)

	result := workers.GetMetaInfo(*filename)
	workers.PrintArray(result)
}
