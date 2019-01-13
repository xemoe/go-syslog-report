package main

import (
	"flag"
	"fmt"
	parser "github.com/xemoe/go-syslog-report/parser"
	"log"
	"strings"
)

var (
	filename = flag.String("f", "syslog.log.gz", "the file to process")
)

func main() {

	flag.Parse()
	fmt.Errorf("Read input: %q\n", *filename)

	files := strings.Split(*filename, ",")
	_, err := parser.ValidateFilesIsExists(files)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Errorf("Read files: %q\n", files)

	result := parser.CountMultiplesSyslogLines(files)
	fmt.Println(result)
}
