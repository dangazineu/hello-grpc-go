package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"dgzn.dev/hello/api"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	api.UnimplementedGreeterServer
}

func (s *server) SayHelloOnce(_ context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &api.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
