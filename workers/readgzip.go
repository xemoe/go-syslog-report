package workers

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	mapper "github.com/xemoe/go-syslog-report/mapper"
	types "github.com/xemoe/go-syslog-report/types"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func PrintArray(result interface{}) {
	p, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(p))
}

func GetDefaultIndex() types.SyslogLineIndex {
	return types.SyslogLineIndex{
		Ts:       0,
		Uid:      1,
		Orig_h:   2,
		Orig_p:   3,
		Resp_h:   4,
		Resp_p:   5,
		Proto:    6,
		Facility: 7,
		Severity: 8,
		Message:  9,
	}
}

func GetSyslogLines(filename string, index types.SyslogLineIndex) *types.SyslogLines {

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
	var result = types.SyslogLines{Index: index}
	ref := reflect.ValueOf(index)
	numfields := ref.NumField()

	for scanner.Scan() {
		line := scanner.Text()
		item, isok := mapper.SyslogMapper(ref, numfields, line)
		if isok {
			result.Contents = append(
				result.Contents,
				item)
		}
	}

	return &result
}

func GetMetaInfo(filename string) *types.MetaInfo {

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
	result := new(types.MetaInfo)

	scanner.Split(bufio.ScanLines)
	rxpmeta := regexp.MustCompile(`^#([^\s]+)\s+(.+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if rxpmeta.MatchString(line) {

			match := rxpmeta.FindStringSubmatch(line)
			key := strings.TrimSpace(match[1])
			value := strings.TrimSpace(match[2])

			switch key {
			case "separator":
				result.Separator = value
			case "set_separator":
				result.Set_separator = value
			case "empty_field":
				result.Empty_field = value
			case "unset_field":
				result.Unset_field = value
			case "path":
				result.Path = value
			case "open":
				result.Open = value
			case "fields":
				result.Fields = strings.Split(value, "\t")
			case "types":
				result.Types = strings.Split(value, "\t")
			}
		} else {
			break
		}
	}

	return result
}
