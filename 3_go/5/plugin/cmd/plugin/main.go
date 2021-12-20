package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/go-plugin/examples/grpc/shared"
	custom "plugin/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"greeter": custom.NewGreeterPlugin(),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
