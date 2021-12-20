package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"os"
	"os/exec"
	"tt_go_plugin_server/api"
)

func main() {

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			"greeter": &api.GreeterPlugin{},
		},
		Cmd:              exec.Command("/Users/dgodyna/work/My/tt_plugins/3_go/5/soulution/plugin"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	kv := raw.(api.GreeterServer)
	fmt.Println(kv.Greet(context.Background(), &api.Request{Name: "test"}))
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
