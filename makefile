#dir = $(shell pwd)
#GOPATH=$(shell go env GOPATH)

prepare:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1
lint:
	cd ./bin ; \
	./golangci-lint run ../...
