package workers

import (
	"bufio"
	gzip "github.com/klauspost/pgzip"
	"log"
	"os"
)

func CountMultiplesSyslogLines(files []string) int {
	sum := 0
	results := make(chan int, 100)
	for i := 0; i < len(files); i++ {
		go func(filename string, results chan int) {
			results <- CountSyslogLines(filename)
		}(files[i], results)
		sum += <-results
	}

	return sum
}

func CountSyslogLines(filename string) int {

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
	counter := 0

	//
	// Skip meta for 8 lines
	//
	skip := 8
	for scanner.Scan() {
		counter = counter + 1
	}

	return counter - skip
}
