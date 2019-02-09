module workers

require (
	github.com/AndreasBriese/bbloom v0.0.0-20180913140656-343706a395b7 // indirect
	github.com/dgraph-io/badger v1.5.4
	github.com/dgryski/go-farm v0.0.0-20190104051053-3adb47b1fb0f // indirect
	github.com/go-test/deep v1.0.1
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/klauspost/compress v1.4.1 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/pgzip v1.2.1
	github.com/xemoe/go-syslog-report/cache v0.0.0
	github.com/xemoe/go-syslog-report/debug v0.0.0
	github.com/xemoe/go-syslog-report/input v0.0.0
	github.com/xemoe/go-syslog-report/mapper v0.0.0
	github.com/xemoe/go-syslog-report/types v0.0.0
	github.com/xemoe/go-syslog-report/validators v0.0.0
	golang.org/x/net v0.0.0-20190206173232-65e2d4e15006 // indirect
	golang.org/x/sys v0.0.0-20190204203706-41f3e6584952 // indirect
)

replace github.com/xemoe/go-syslog-report/types => ../types

replace github.com/xemoe/go-syslog-report/mapper => ../mapper

replace github.com/xemoe/go-syslog-report/validators => ../validators

replace github.com/xemoe/go-syslog-report/input => ../input

replace github.com/xemoe/go-syslog-report/debug => ../debug

replace github.com/xemoe/go-syslog-report/cache => ../cache
