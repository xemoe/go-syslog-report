package parser

import (
	"reflect"
	"strings"
)

var syslog = SyslogLine{}
var fType = reflect.TypeOf(syslog)
var fVal = reflect.New(fType)

func SyslogMapper(ref reflect.Value, numfields int, line string) (SyslogLine, bool) {
	if !strings.HasPrefix(line, "#") {
		values := strings.Split(line, "\t")
		for i := 0; i < numfields; i++ {
			idx := ref.Field(i).Interface().(int)
			if idx >= 0 {
				ival := strings.TrimSpace(values[idx])
				fVal.Elem().Field(idx).SetString(ival)
			}
		}
		return fVal.Elem().Interface().(SyslogLine), true
	}

	return SyslogLine{}, false
}

func SyslogGroupMapper(ref reflect.Value, numfields int, line string) (string, bool) {
	if !strings.HasPrefix(line, "#") {
		uniqstr := []string{}
		values := strings.Split(line, "\t")

		for i := 0; i < numfields; i++ {
			idx := ref.Field(i).Interface().(int)
			if idx >= 0 {
				ival := strings.TrimSpace(values[idx])
				uniqstr = append(uniqstr, ival)
			}
		}

		return strings.Join(uniqstr, "|"), true
	}

	return "", false
}
