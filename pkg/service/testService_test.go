package service

import (
	"context"
	"net"
	"testing"

	"github.com/cwza/circleci-test/pkg/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestPing(t *testing.T) {
	serverListener, server := newGrpcServer(t)
	clientConn, client := newGrpcClient(t, serverListener.Addr().String())
	defer closeAll(t, clientConn, server, serverListener)

	value := "testtest"
	expected := proto.PingResponse{Value: &value}
	res, err := client.Ping(context.Background(), &proto.PingRequest{Value: &value})
	require.NoError(t, err, "should not have error while Ping")
	assert.Equal(t, expected, *res)
}

func newGrpcServer(t *testing.T) (net.Listener, *grpc.Server) {
	////// Grpc Server
	grpcLis, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "must be able to allocate a port for serverListener")
	grpcServer := grpc.NewServer()

	proto.RegisterTestServiceServer(grpcServer, &TestService{})

	//////// Serve
	go func() {
		err := grpcServer.Serve(grpcLis)
		require.NoError(t, err, "must not error while grpc server serve")
	}()
	return grpcLis, grpcServer
}

func newGrpcClient(t *testing.T, serverAddr string) (*grpc.ClientConn, proto.TestServiceClient) {
	clientConn, err := grpc.Dial(
		serverAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	require.NoError(t, err, "must not error on client Dial")

	client := proto.NewTestServiceClient(clientConn)

	return clientConn, client
}

func closeAll(t *testing.T, clientConn *grpc.ClientConn, server *grpc.Server, serverListener net.Listener) {
	clientConn.Close()
	server.Stop()
	serverListener.Close()
}
