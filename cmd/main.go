package main

import (
	mygrpc "github.com/imran4u/go-grpc-unary-server/internal/adapter/grpc"
	app "github.com/imran4u/go-grpc-unary-server/internal/application"
)

func main() {
	// mygrpc.NewGrpcAdapter(app.NewHelloService(), 8080).Run()
	hs := &app.HelloService{}
	grpcAdapter := mygrpc.NewGrpcAdapter(hs, 8080)
	grpcAdapter.Run()
}
