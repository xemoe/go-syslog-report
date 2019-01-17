package input

import (
	deep "github.com/go-test/deep"
	types "github.com/xemoe/go-syslog-report/types"
	"testing"
)

func TestParseIndexFields_withEmptyFields_ShouldReturnExpectedValues(t *testing.T) {

	expected := types.SyslogLineIndex{
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

	input := ""
	result := ParseIndexFields(input)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestParseIndexFields_withSelectedFields_ShouldReturnExpectedValues(t *testing.T) {

	expected := types.SyslogLineIndex{
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

	input := "id.orig_h,facility"
	result := ParseIndexFields(input)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestParseIndexFields_withSomeInvalidFields_ShouldReturnExpectedValues(t *testing.T) {

	expected := types.SyslogLineIndex{
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

	input := "id.orig_h,facility,foo"
	result := ParseIndexFields(input)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}
