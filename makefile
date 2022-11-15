#dir = $(shell pwd)
#GOPATH=$(shell go env GOPATH)

prepare:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1

lint:
	cd ./bin ; \
	./golangci-lint run ../...
up :
	cd ./build ;\
	docker compose up
down:
	cd ./build ;\
	docker compose down

start:
	cd ./build ;\
	docker compose start

stop:
	cd ./build ;\
	docker compose stop