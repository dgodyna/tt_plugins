package main

import (
	"fmt"
	"github.com/fatih/color"
	"plugin"
	"tt_go_plugin_server/pkg/api"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(fmt.Errorf("unable to open plugin : %w", err))
	}
	lookup, err := p.Lookup("ProvideEcher")
	if err != nil {
		panic(err)
	}

	var provider = lookup.(func() api.Echer)
	echer := provider()
	color.Green("Calling Pugin")
	echer.Echo("gophers")
}
