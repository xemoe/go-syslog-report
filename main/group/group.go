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

	//
	// @TODO
	// to use with flag
	// -g id.orig_h,Facility
	//
	skipIndex := parser.SyslogLineIndex{
		Ts:       -1,
		Uid:      -1,
		Orig_h:   2,
		Orig_p:   -1,
		Resp_h:   -1,
		Resp_p:   -1,
		Proto:    -1,
		Facility: 7,
		Severity: -1,
		Message:  -1,
	}

	result := parser.GroupCountMultiples(files, skipIndex)
	for _, v := range result {
		fmt.Printf("%s;%d\n", v.Key, v.Value)
	}
}
