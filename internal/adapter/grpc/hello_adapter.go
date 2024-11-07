package grpc

import (
	"context"

	"github.com/imran4u/go-grpc-proto/protogen/go/hello"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {

	return &hello.HelloResponse{
		Greet: a.helloservice.GenerateHello(req.Name),
	}, nil
}
