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
	_ = NewTest()
	//	just empty main file to build pluginn
}

type test struct {
	s string
}

//export NewTest
//go:noinline
func NewTest() *test {
	n := GetName()
	return &test{
		s: n,
	}
}

func GetName() string {

	return "tetete"
}
