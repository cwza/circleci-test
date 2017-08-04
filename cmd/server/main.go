package main

import (
	"log"
	"net"
	"os"

	"github.com/cwza/circleci-test/pkg/proto"
	"github.com/cwza/circleci-test/pkg/service"
	"google.golang.org/grpc"
)

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
	proto.RegisterTestServiceServer(grpcServer, &service.TestService{})

	if err := grpcServer.Serve(grpcLis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
