package server

import (
    "context"
	
    "github.com/t3mp14r3/curly-octopus/checks/gen"
)

func (c *Service) Create(ctx context.Context, req *gen.CreateRequest) (*gen.CreateResponse, error) {
    return &gen.CreateResponse{
        Filename: "hello ",
        Data: []byte("world"),
    }, nil
}
