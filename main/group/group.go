package main

import (
	"flag"
	"fmt"
	types "github.com/xemoe/go-syslog-report/types"
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
	log.Printf("Read syslog from input:%s", *filename)

	files := strings.Split(*filename, ",")
	_, err := validators.ValidateFilesIsExists(files)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Read syslog from files:%q", files)

	//
	// @TODO
	// to use with flag
	// -g id.orig_h,Facility
	//
	skipIndex := types.SyslogLineIndex{
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

	result := workers.GroupCountMultiples(files, skipIndex)
	for _, v := range result {
		fmt.Printf("%s;%d\n", v.Key, v.Value)
	}
}
