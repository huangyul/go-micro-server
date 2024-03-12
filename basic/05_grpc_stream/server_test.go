package _5_grpc_stream

import (
	"fmt"
	"go-micro-server/basic/05_grpc_stream/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"testing"
	"time"
)

type Server struct {
	proto.UnimplementedHelloServer
}

func (s Server) GetStream(request *proto.Request, server proto.Hello_GetStreamServer) error {
	i := 0
	for i < 3 {
		_ = server.Send(&proto.Response{
			Data: fmt.Sprintf("%v", time.Now().UnixMilli()),
		})
		time.Sleep(time.Second * 2)
		i++
	}
	return nil
}

func (s Server) PostStream(server proto.Hello_PostStreamServer) error {
	for {
		if v, err := server.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(v)
		}
	}
	return nil
}

func (s Server) AllStream(server proto.Hello_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := server.Recv()
			fmt.Println("接受到客户端的信息" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = server.Send(&proto.Response{
				Data: "我是服务端",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

	return nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic("listen失败")
	}
	g := grpc.NewServer()
	proto.RegisterHelloServer(g, &Server{})
	err = g.Serve(lis)
	if err != nil {
		panic("启动失败")
	}
}
