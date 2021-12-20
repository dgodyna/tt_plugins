package main

import (
	"plugin/pkg/echo"
	"tt_go_plugin_server/pkg/api"
)

// ProvideEcher create a new instance of api.Echer.
func ProvideEcher() api.Echer {
	return echo.NewEcher()
}

func main() {
	//	just empty main file to build pluginn
}
