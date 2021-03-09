package adder

import (
	"context"
	"grpc_my_example/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

// Add ...
func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}