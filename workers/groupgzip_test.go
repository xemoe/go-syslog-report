package workers

import (
	deep "github.com/go-test/deep"
	types "github.com/xemoe/go-syslog-report/types"
	"testing"
)

func TestGroupCount_WithSkipIndex_ShouldReturnExpectedValues(t *testing.T) {

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

	expected := map[string]int{
		"172.16.2.168|KERN": 2,
	}

	filename := "files/test1.log.gz"
	result := GroupCount(filename, skipIndex)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestGroupCountMultiples_WithSkipIndex_ShouldReturnExpectedValues(t *testing.T) {

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

	expected := []types.GroupCountSlices{
		types.GroupCountSlices{"172.16.2.168|KERN", 5},
	}

	files := []string{
		"files/test1.log.gz",
		"files/test2.log.gz",
	}
	result := GroupCountMultiples(files, skipIndex)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}
