package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 1. 初始化一个 server
	listener, _ := net.Listen("tcp", ":8088")
	// 2. 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloServer{})
	// 3. 启动服务
	conn, _ := listener.Accept()
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}
