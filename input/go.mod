module input

require (
	github.com/go-test/deep v1.0.1
	github.com/xemoe/go-syslog-report/debug v0.0.0
	github.com/xemoe/go-syslog-report/types v0.0.0
)

replace github.com/xemoe/go-syslog-report/types => ../types

replace github.com/xemoe/go-syslog-report/debug => ../debug
