package workers

import (
	deep "github.com/go-test/deep"
	"testing"
)

func TestCountSyslogLines_ShouldReturnExpectedValues(t *testing.T) {

	expected := 2

	filename := "files/test1.log.gz"
	result := CountSyslogLines(filename)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestCountMultiplesSyslogLines_ShouldReturnExpectedValues(t *testing.T) {

	expected := 5

	files := []string{
		"files/test1.log.gz",
		"files/test2.log.gz",
	}
	result := CountMultiplesSyslogLines(files)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}
