package main

import (
	"log"
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

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("连接rpc服务失败", err)
	}

	//release
	defer client.Close()

	var (
		req  = GetUserReq{Id: "1"}
		resp GetUserResp
	)

	err = client.Call("UserServer.GetUser", req, &resp)
	if err != nil {
		log.Fatal("调用rpc服务失败", err)
	}

	log.Println(resp)

}
