proto:
	protoc --go_out=plugins=grpc:. ./pkg/proto/*.proto