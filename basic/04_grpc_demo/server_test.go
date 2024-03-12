package grpc_demo_test

import (
	"context"
	"go-micro-server/basic/04_grpc_demo/proto"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type Server struct {
	proto.UnimplementedHelloServer
}

func (s Server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.Response, error) {
	return &proto.Response{
		Reply: "hello," + request.Name,
	}, nil
}
func TestGrpc(t *testing.T) {
	// 新建grpc服务
	g := grpc.NewServer()
	// 注册方法
	proto.RegisterHelloServer(g, &Server{})
	// 启动服务
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic("启动失败")
	}
	err = g.Serve(lis)
	if err != nil {
		panic("启动失败2")
	}
}
