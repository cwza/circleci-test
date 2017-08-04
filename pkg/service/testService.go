package service

import (
	"github.com/cwza/circleci-test/pkg/proto"
	"golang.org/x/net/context"
)

type TestService struct{}

func (s *TestService) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	value := req.GetValue()
	return &proto.PingResponse{Value: &value}, nil
}
