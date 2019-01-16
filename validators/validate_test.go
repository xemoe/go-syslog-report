package validators

import (
	deep "github.com/go-test/deep"
	"testing"
)

func TestValidateFilesIsExists_WhenExists(t *testing.T) {

	expected := true

	result, _ := ValidateFilesIsExists([]string{
		"files/test1.log.gz",
		"files/test2.log.gz",
	})

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestValidateFilesIsExists_WhenNotExists(t *testing.T) {

	expected := false

	result, _ := ValidateFilesIsExists([]string{
		"files/test1.log.gz",
		"files/test2.log.gz",
		"files/test3.log.gz",
		"files/test4.log.gz",
	})

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}
