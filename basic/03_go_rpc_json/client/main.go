package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		panic("连接失败")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	reply := ""
	client.Call("HelloService.Hello", "test", &reply)

	fmt.Println(reply)
}
