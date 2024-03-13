package grpc_demo_test

import (
	"context"
	"fmt"
	"go-micro-server/basic/04_grpc_demo/proto"
	"testing"

	"google.golang.org/grpc"
)

func TestGrpcCilent(t *testing.T) {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	c := proto.NewHelloClient(conn)
	r, err := c.Hello(context.Background(), &proto.HelloRequest{Name: "grpc"})
	if err != nil {
		panic("调用失败")
	}
	fmt.Printf(r.Reply)
}
