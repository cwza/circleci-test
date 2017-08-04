package main

import (
	"log"
	"os"

	"github.com/cwza/circleci-test/pkg/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
	}
	conn, err := grpc.Dial(
		"127.0.0.1:"+port,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("can not connect: %+v", err)
	}
	defer conn.Close()

	client := proto.NewTestServiceClient(conn)
	value := "testtest"
	res, err := client.Ping(context.Background(), &proto.PingRequest{Value: &value})
	if err != nil {
		log.Fatalf("can not get res err: %+v", err)
	}
	log.Printf("res: %+v", res)
}
