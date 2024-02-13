package main

import (
	"context"
	"fmt"
	"go-micro-server/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "huang"})

	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}
