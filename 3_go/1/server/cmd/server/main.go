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

	var provider, ok = lookup.(func() api.Echer)
	if !ok {
		panic(fmt.Errorf("unabe to cust loaded function 'ProvideEcher' with type '%T' to 'func() api.Echer'", lookup))
	}
	echer := provider()
	color.Green("Calling Pugin")
	echer.Echo("gophers")
}
