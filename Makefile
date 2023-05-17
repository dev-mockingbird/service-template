VERSION ?= $(shell git describe --always --tags)
HUB ?= localhost:5001/mockingbird/channel
IMAGE ?= ${HUB}:$(VERSION)
.PHONY: init
init:
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: update
update:
	@go get -u

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
