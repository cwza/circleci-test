TEST_PACKAGES := $(shell go list ./... | grep -vE '(vendor|cmd|mock|proto|integrationtest)')

test-unit:
	echo "test-unit"

test:
	echo "test"

test-integration:
	echo "test-integration"