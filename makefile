#GOPATH=$(shell go env GOPATH)
rootDir = $(shell pwd)
dc = docker compose -f ${rootDir}/build/docker-compose.yaml --env-file ${rootDir}/build/.local.env
prepare:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1
	go install github.com/golang/mock/mockgen@v1.6.0

lint:
	cd ./bin ; \
	./golangci-lint run ../...

up:
	${dc} up

down:
	${dc} down

build:
	${dc} build

create:
	migrate create -ext sql -dir db/migration -seq $(name)
