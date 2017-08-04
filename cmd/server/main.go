package main

import (
	"log"
	"net"
	"os"

	"github.com/cwza/circleci-test/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type TestService struct{}

func (s *TestService) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	value := req.GetValue()
	return &proto.PingResponse{Value: &value}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
	}

	grpcLis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen for grpc: %+v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTestServiceServer(grpcServer, &TestService{})

	if err := grpcServer.Serve(grpcLis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
