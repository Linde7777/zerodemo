package logic

import (
	"context"

	"zerodemo/user/rpc/internal/svc"
	"zerodemo/user/rpc/pb/user"

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

func (l *GetUserLogic) GetUser(in *user.ReqGetUser) (*user.RespGetUser, error) {
	// todo: add your logic here and delete this line

	return &user.RespGetUser{}, nil
}
