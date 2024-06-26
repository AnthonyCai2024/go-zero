package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero/user/rpc/models"
	"gozero/user/rpc/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
