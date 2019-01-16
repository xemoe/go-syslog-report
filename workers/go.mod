module workers

require (
	github.com/go-test/deep v1.0.1
	github.com/klauspost/compress v1.4.1 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/pgzip v1.2.1
	github.com/xemoe/go-syslog-report/mapper v0.0.0
	github.com/xemoe/go-syslog-report/types v0.0.0
	github.com/xemoe/go-syslog-report/validators v0.0.0
)

replace github.com/xemoe/go-syslog-report/types => ../types

replace github.com/xemoe/go-syslog-report/mapper => ../mapper

replace github.com/xemoe/go-syslog-report/validators => ../validators
