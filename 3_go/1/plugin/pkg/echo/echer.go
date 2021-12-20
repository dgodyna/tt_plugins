package echo

import (
	"fmt"
	"tt_go_plugin_server/pkg/api"
)

var _ api.Echer = (*echer)(nil)

// echer is realization of api.Echer which will just print hello to stdout.
type echer struct {
}

func (e *echer) Echo(s string) {
	fmt.Printf("hello '%s' from custom plugin\n", s)
}

// NewEcher create a new instance of api.Echer which just print hello to stdout.
func NewEcher() api.Echer {
	return &echer{}
}
