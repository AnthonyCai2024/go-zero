package main

import (
	"context"
	"errors"
	"go-zero/grpc/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServer struct{}

func (*UserServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	if u, ok := users[req.Id]; ok {
		return &user.GetUserResponse{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return nil, errors.New("没有找到用户")

}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, new(UserServer))

	log.Println("服务启动成功")

	err = s.Serve(listen)
}
