package mapper

import (
	types "github.com/xemoe/go-syslog-report/types"
	"reflect"
	"strings"
)

func SyslogMapper(ref reflect.Value, numfields int, line string) (types.SyslogLine, bool) {
	if !strings.HasPrefix(line, "#") {

		syslog := types.SyslogLine{}
		fType := reflect.TypeOf(syslog)
		fVal := reflect.New(fType)

		values := strings.Split(line, "\t")
		for i := 0; i < numfields; i++ {
			idx := ref.Field(i).Interface().(int)
			if idx >= 0 {
				ival := strings.TrimSpace(values[idx])
				fVal.Elem().Field(idx).SetString(ival)
			}
		}
		return fVal.Elem().Interface().(types.SyslogLine), true
	}

	return types.SyslogLine{}, false
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
