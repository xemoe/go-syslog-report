package workers

import (
	deep "github.com/go-test/deep"
	types "github.com/xemoe/go-syslog-report/types"
	"testing"
)

func TestGetMetaInfo_ShouldReturnExpectedValues(t *testing.T) {

	expected := &types.MetaInfo{
		Separator:     "\\x09",
		Set_separator: ",",
		Empty_field:   "(empty)",
		Unset_field:   "-",
		Path:          "syslog",
		Open:          "2019-01-08-10-00-00",
		Fields: []string{
			"ts", "uid", "id.orig_h", "id.orig_p", "id.resp_h", "id.resp_p", "proto", "facility", "severity", "message",
		},
		Types: []string{
			"time", "string", "addr", "port", "addr", "port", "enum", "string", "string", "string",
		},
	}

	filename := "files/test1.log.gz"
	result := GetMetaInfo(filename)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}
