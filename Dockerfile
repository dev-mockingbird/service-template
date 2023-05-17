FROM golang:1.19.0-alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY . .
RUN make tidy build

# Deployment container
FROM alpine

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/build/channel /channel

ENTRYPOINT [ "/channel" ]
