package api

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GreeterPlugin struct {
	plugin.Plugin
	Impl GreeterServer
}

func (p *GreeterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterGreeterServer(s, p.Impl)
	return nil
}

func (p *GreeterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewGreeterClient(c)}, nil
}

var _ GreeterServer = (*GRPCClient)(nil)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client GreeterClient }

func (g *GRPCClient) mustEmbedUnimplementedGreeterServer() {
	panic("implement me")
}

func (g *GRPCClient) Greet(ctx context.Context, request *Request) (*Response, error) {
	return g.client.Greet(ctx, request)
}
