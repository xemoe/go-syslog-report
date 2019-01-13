package parser

import (
	"github.com/go-test/deep"
	"testing"
)

func TestGetMetaInfo_ShouldReturnExpectedValues(t *testing.T) {

	expected := &MetaInfo{
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

func TestGetSyslogLines_WithDefaultIndex_ShouldReturnExpectedValues(t *testing.T) {

	expected := &SyslogLines{
		Index: GetDefaultIndex(),
		Contents: []SyslogLine{
			SyslogLine{
				"1546916399.834895",
				"COxRoS1BBRMVHBi0u3",
				"172.16.2.168",
				"60915",
				"172.16.2.181",
				"514",
				"udp",
				"KERN",
				"INFO",
				`"2019-01-08 10:00:01","2694336246","172.29.3.201","Packet filter","Notification","New connection","Allow","172.29.7.201","172.217.24.174","HTTPS","6","59056","443","158.0",,,,,"948436992","Interface #1",,,,,,,,,,,,,,"SGFW node 1",,,,,"2019-01-08 10:00:01","Firewall","Connection_Allowed",,,"6488237649027419894",,,,`,
			},
			SyslogLine{
				"1546916399.834895",
				"COxRoS1BBRMVHBi0u3",
				"172.16.2.168",
				"60915",
				"172.16.2.181",
				"514",
				"udp",
				"KERN",
				"INFO",
				`"2019-01-08 10:00:01","2694336247","172.29.3.201","Packet filter","Notification","New connection","Allow","172.29.7.201","172.217.24.174","HTTPS","6","59056","443","158.0",,,,,"948436992","Interface #1",,,,,,,,,,,,,,"SGFW node 1",,,,,"2019-01-08 10:00:01","Firewall","Connection_Allowed",,,"6488237649027419895",,,,`,
			},
		},
	}

	filename := "files/test1.log.gz"
	result := GetSyslogLines(filename, GetDefaultIndex())

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

func TestGetSyslogLines_WithSkipIndex_ShouldReturnExpectedValues(t *testing.T) {

	skipIndex := SyslogLineIndex{
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

	expected := &SyslogLines{
		Index: skipIndex,
		Contents: []SyslogLine{
			SyslogLine{
				Orig_h:   "172.16.2.168",
				Facility: "KERN",
			},
			SyslogLine{
				Orig_h:   "172.16.2.168",
				Facility: "KERN",
			},
		},
	}

	filename := "files/test1.log.gz"
	result := GetSyslogLines(filename, skipIndex)

	if diff := deep.Equal(expected, result); diff != nil {
		t.Error(diff)
	}
}

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

func TestGroupCount_WithSkipIndex_ShouldReturnExpectedValues(t *testing.T) {

	skipIndex := SyslogLineIndex{
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

	skipIndex := SyslogLineIndex{
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

	expected := []GroupCountSlices{
		GroupCountSlices{"172.16.2.168|KERN", 5},
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
