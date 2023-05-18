VERSION ?= $(shell git describe --always --tags)
HUB ?= localhost:5001/mockingbird/channel
IMAGE ?= ${HUB}:$(VERSION)
.PHONY: init
init:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: update
update:
	@go get -u

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc/proto/helloworld.proto

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build: 
	@go build -o build/channel main.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: swagger
swagger:
	@swag init --dir http --output http/docs --generalInfo handler.go

.PHONY: docker
docker:
	@docker build -t ${IMAGE} .
