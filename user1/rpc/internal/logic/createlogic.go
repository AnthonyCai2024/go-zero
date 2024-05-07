package logic

import (
	"context"

	"cai/gozero/user/rpc/internal/svc"
	"cai/gozero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line

	return &user.CreateResp{}, nil
}
