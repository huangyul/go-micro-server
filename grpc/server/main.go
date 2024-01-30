package server

import (
	"context"
	"go-micro-server/grpc/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello" + req.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()

	proto.RegisterGreeterServer(g, &Server{})
}
