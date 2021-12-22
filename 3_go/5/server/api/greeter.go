package api

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var _ plugin.GRPCPlugin = (*GreeterPlugin)(nil)

type GreeterPlugin struct {
	plugin.Plugin
	Impl GreeterServer
}

// GRPCServer should register this plugin for serving with the
// given GRPCServer. Unlike Plugin.Server, this is only called once
// since gRPC plugins serve singletons.
func (p *GreeterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterGreeterServer(s, p.Impl)
	return nil
}

// GRPCClient should return the interface implementation for the plugin
// you're serving via gRPC. The provided context will be canceled by
// go-plugin in the event of the plugin process exiting.
func (p *GreeterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewGreeterClient(c)}, nil
}

var _ GreeterServer = (*GRPCClient)(nil)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client GreeterClient }

// this is a bad example, please create own interface and don't use existing gRPC server
func (g *GRPCClient) mustEmbedUnimplementedGreeterServer() {
	panic("implement me")
}

func (g *GRPCClient) Greet(ctx context.Context, request *Request) (*Response, error) {
	return g.client.Greet(ctx, request)
}
