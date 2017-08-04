TEST_PACKAGES := $(shell go list ./... | grep -vE '(vendor|cmd|mock|proto|integrationtest)')

test:
	go test ${TEST_PACKAGES}

proto:
	protoc --go_out=plugins=grpc:. ./pkg/proto/*.proto