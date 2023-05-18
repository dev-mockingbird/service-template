package grpc

import (
	"context"

	pb "udious.com/mockingbird/channel/grpc/proto"
)

type Handler struct {
	pb.UnimplementedGreeterServer
}

func (h Handler) Sayhello(ctx context.Context, in *pb.Hellorequest) (*pb.Helloreply, error) {
	return &pb.Helloreply{Message: "Hello " + in.GetName()}, nil
}
