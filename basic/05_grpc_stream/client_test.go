package _5_grpc_stream

import (
	"context"
	"fmt"
	"go-micro-server/basic/05_grpc_stream/proto"
	"google.golang.org/grpc"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := proto.NewHelloClient(conn)

	// 客服端流模式
	//res, _ := c.GetStream(context.Background(), &proto.Request{
	//	Data: "hello",
	//})
	//for {
	//	a, err := res.Recv()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(a)
	//}

	// 服务端流模式
	//postS, err := c.PostStream(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//i := 0
	//for i < 3 {
	//	_ = postS.Send(&proto.Request{Data: "client"})
	//	time.Sleep(time.Second)
	//	i++
	//}

	// 双向流模式
	client, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			_ = client.Send(&proto.Request{
				Data: "来自客户端发送的信息",
			})
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			res, _ := client.Recv()
			fmt.Println(res.Data)
		}
	}()

	wg.Wait()
}
