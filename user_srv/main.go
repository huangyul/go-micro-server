package main

import (
	"flag"
	"fmt"
	"go-micro-server/user_srv/handler"
	"go-micro-server/user_srv/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 通过flag获取参数
	ip := flag.String("ip", "0.0.0.0", "ip地址")
	port := flag.String("port", "8088", "端口号")
	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%s`, *ip, *port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
