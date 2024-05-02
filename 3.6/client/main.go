package main

import (
	"context"
	"go-zero/grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	client, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("连接rpc服务失败", err)
	}

	//release
	defer client.Close()

	c := user.NewUserServiceClient(client)

	res, err := c.GetUser(context.Background(), &user.GetUserRequest{Id: "1"})

	if err != nil {
		log.Fatal("调用rpc服务失败", err)
	}

	log.Println(res)
}
