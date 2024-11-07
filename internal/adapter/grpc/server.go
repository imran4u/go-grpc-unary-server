package grpc

import (
	"net"
	"strconv"

	"github.com/imran4u/go-grpc-proto/protogen/go/hello"
	"github.com/imran4u/go-grpc-unary-server/internal/port"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloservice port.HelloServicePort
	grpcPort     int
	grpcServer   *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	// grpcServer := grpc.NewServer()
	return &GrpcAdapter{
		helloservice: helloService,
		grpcPort:     grpcPort,
		// grpcServer:   grpcServer,
	}
}

func (s *GrpcAdapter) Run() {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(s.grpcPort))
	if err != nil {
		panic(err)
	}
	println("Listening on port", s.grpcPort)

	grpcServer := grpc.NewServer()
	s.grpcServer = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, s)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (s *GrpcAdapter) Stop() {
	s.grpcServer.Stop()
}
