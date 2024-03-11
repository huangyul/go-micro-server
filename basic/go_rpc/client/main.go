package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8088")
	if err != nil {
		panic("连接失败")
	}
	var reply *string = new(string)
	err = client.Call("HelloService.Hello", "test", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
