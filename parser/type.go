package parser

type MetaInfo struct {
	Separator     string
	Set_separator string
	Empty_field   string
	Unset_field   string
	Path          string
	Open          string
	Fields        []string
	Types         []string
}

type SyslogLine struct {
	Ts       string `bson:",omitempty"`
	Uid      string `bson:",omitempty"`
	Orig_h   string `bson:",omitempty"`
	Orig_p   string `bson:",omitempty"`
	Resp_h   string `bson:",omitempty"`
	Resp_p   string `bson:",omitempty"`
	Proto    string `bson:",omitempty"`
	Facility string `bson:",omitempty"`
	Severity string `bson:",omitempty"`
	Message  string `bson:",omitempty"`
}

type SyslogLineIndex struct {
	Ts       int `bson:",omitempty"`
	Uid      int `bson:",omitempty"`
	Orig_h   int `bson:",omitempty"`
	Orig_p   int `bson:",omitempty"`
	Resp_h   int `bson:",omitempty"`
	Resp_p   int `bson:",omitempty"`
	Proto    int `bson:",omitempty"`
	Facility int `bson:",omitempty"`
	Severity int `bson:",omitempty"`
	Message  int `bson:",omitempty"`
}

type SyslogLines struct {
	Contents []SyslogLine
	Index    SyslogLineIndex
}

type GroupCountSlices struct {
	Key   string
	Value int
}
