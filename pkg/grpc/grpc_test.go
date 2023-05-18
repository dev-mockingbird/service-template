package grpc

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpch "udious.com/mockingbird/channel/grpc"
	pb "udious.com/mockingbird/channel/grpc/proto"
)

func TestGrpc_Start(t *testing.T) {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, grpch.Handler{})
	rs := StartStopper{Server: s, Port: 8000}
	errCh := rs.Start()
	time.Sleep(time.Second)
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := c.Sayhello(ctx, &pb.Hellorequest{Name: "my lord"})
	if err != nil {
		t.Fatal(err)
	}
	if r.Message != "Hello my lord" {
		t.Fatal("received false message")
	}
	rs.Stop(context.Background())
	if err := <-errCh; err != nil {
		t.Fatal("server not stop gracefully")
	}
}
