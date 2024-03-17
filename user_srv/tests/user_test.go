package tests

import (
	"context"
	"fmt"
	"go-micro-server/user_srv/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func newUserClient() proto.UserClient {
	conn, err := grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		panic(err)
	}
	client := proto.NewUserClient(conn)
	return client
}

func TestCreateUser(t *testing.T) {
	client := newUserClient()
	user, err := client.CreateUser(context.Background(), &proto.CreateUserRequest{
		Mobile:   "123123123",
		NickName: "test",
		Password: "123123",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
