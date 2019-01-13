package main

import (
	"fmt"
)

func main() {
	files := []string{
		"large1.log.gz",
		"large2.log.gz",
		"large3.log.gz",
		"large4.log.gz",
	}
	result := CountMultiplesSyslogLines(files)
	fmt.Println(result)
}
