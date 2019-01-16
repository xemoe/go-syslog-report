package workers

import (
	deep "github.com/go-test/deep"
	types "github.com/xemoe/go-syslog-report/types"
	"testing"
)

func TestGetSyslogLines_WithDefaultIndex_ShouldReturnExpectedValues(t *testing.T) {

	expected := &types.SyslogLines{
		Index: GetDefaultIndex(),
		Contents: []types.SyslogLine{
			types.SyslogLine{
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
			types.SyslogLine{
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

	expected := &types.SyslogLines{
		Index: skipIndex,
		Contents: []types.SyslogLine{
			types.SyslogLine{
				Orig_h:   "172.16.2.168",
				Facility: "KERN",
			},
			types.SyslogLine{
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
