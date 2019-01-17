module main

require (
	github.com/xemoe/go-syslog-report/debug v0.0.0
	github.com/xemoe/go-syslog-report/input v0.0.0
	github.com/xemoe/go-syslog-report/mapper v0.0.0
	github.com/xemoe/go-syslog-report/types v0.0.0
	github.com/xemoe/go-syslog-report/validators v0.0.0
	github.com/xemoe/go-syslog-report/workers v0.0.0
)

replace github.com/xemoe/go-syslog-report/workers => ../workers

replace github.com/xemoe/go-syslog-report/types => ../types

replace github.com/xemoe/go-syslog-report/mapper => ../mapper

replace github.com/xemoe/go-syslog-report/validators => ../validators

replace github.com/xemoe/go-syslog-report/input => ../input

replace github.com/xemoe/go-syslog-report/debug => ../debug
