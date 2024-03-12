package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8088")
	rpc.RegisterName("HelloService", &HelloService{})

	for {
		conn, _ := lis.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
