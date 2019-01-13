package main

import (
	"flag"
	parser "github.com/xemoe/go-syslog-report/parser"
)

var (
	filename = flag.String("f", "syslog.log.gz", "the file to process")
)

func main() {
	flag.Parse()
	result := parser.GetMetaInfo(*filename)
	parser.PrintArray(result)
}
