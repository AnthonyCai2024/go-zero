package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozero/user/rpc/api/internal/config"
	"gozero/user/rpc/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}