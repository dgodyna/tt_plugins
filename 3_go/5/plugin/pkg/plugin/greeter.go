package plugin

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"tt_go_plugin_server/api"
)

type greeter struct {
	api.UnimplementedGreeterServer
}

func (g *greeter) Greet(_ context.Context, r *api.Request) (*api.Response, error) {
	return &api.Response{Result: []byte(fmt.Sprintf("hello %s from custom plugin!!!!!1", r.Name))}, nil
}

func NewGreeter() api.GreeterServer {
	return &greeter{}
}

func NewGreeterPlugin() plugin.Plugin {
	return &api.GreeterPlugin{
		Impl: NewGreeter(),
	}
}
