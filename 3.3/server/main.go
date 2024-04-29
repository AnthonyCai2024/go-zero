/**
 * @author: caiqian
 * @doc:
 */

package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type (
	GetUserReq struct {
		Id string `json:"id"`
	}

	GetUserResp struct {
		Id    string
		Name  string
		Phone string
	}
)

type UserServer struct{}

// GetUser UserServer.GetUser
func (*UserServer) GetUser(req GetUserReq, resp *GetUserResp) error {
	// get user for req

	if u, ok := users[req.Id]; ok {
		*resp = GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}

	return errors.New("没有找到用户")
}

func main() {
	// 创建好服务
	userServer := new(UserServer)

	// 服务注册到rpc里
	rpc.Register(userServer)

	// 监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}

	log.Println("服务启动成功")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接收客户端连接失败", err)
			continue
		}

		// 并发处理客户端请求
		go rpc.ServeConn(conn)
	}

}
