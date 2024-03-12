package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8088")

	_ = rpc.RegisterName("HelloService", &HelloService{})

	conn, _ := lis.Accept()
	rpc.ServeConn(conn)

}
