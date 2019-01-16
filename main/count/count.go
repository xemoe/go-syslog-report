package main

import (
	"flag"
	"fmt"
	validators "github.com/xemoe/go-syslog-report/validators"
	workers "github.com/xemoe/go-syslog-report/workers"
	"log"
	"strings"
)

var (
	filename = flag.String("f", "syslog.log.gz", "the file to process")
)

func main() {

	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Read syslog from input:%s", *filename)

	files := strings.Split(*filename, ",")
	_, err := validators.ValidateFilesIsExists(files)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read syslog from files:%q", files)

	result := workers.CountMultiplesSyslogLines(files)
	fmt.Println(result)
}
