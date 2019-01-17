package input

import (
	types "github.com/xemoe/go-syslog-report/types"
	"reflect"
	"strings"
)

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

func ParseIndexFields(input ...string) types.SyslogLineIndex {

	fields := []string{}
	index := GetDefaultIndex()
	if len(input) > 0 && input[0] != "" {
		fields = strings.Split(input[0], ",")
	} else {
		return index
	}

	keys := []string{
		"ts",
		"uid",
		"id.orig_h",
		"id.orig_p",
		"id.resp_h",
		"id.resp_p",
		"proto",
		"facility",
		"severity",
		"message",
	}

	fIndex := reflect.TypeOf(index)
	fVal := reflect.New(fIndex)

	match := false
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(fields); j++ {
			if keys[i] == fields[j] {
				match = true
			}
		}
		if !match {
			fVal.Elem().Field(i).SetInt(-1)
		} else {
			fVal.Elem().Field(i).SetInt(int64(i))
		}

		match = false
	}

	return fVal.Elem().Interface().(types.SyslogLineIndex)
}
