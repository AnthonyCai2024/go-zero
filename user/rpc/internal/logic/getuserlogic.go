package logic

import (
	"context"

	"gozero/user/rpc/rpc/internal/svc"
	"gozero/user/rpc/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	if u, ok := users[in.Id]; ok {
		return &user.GetUserResponse{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return &user.GetUserResponse{}, nil
}
