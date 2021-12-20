module plugin

go 1.17

require (
	github.com/fatih/color v1.13.0
	tt_go_plugin_server v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

replace github.com/fatih/color => github.com/fatih/color v1.12.0

replace tt_go_plugin_server => ../server
