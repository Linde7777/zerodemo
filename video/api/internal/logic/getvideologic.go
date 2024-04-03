package logic

import (
	"context"
	"zerodemo/user/rpc/pb/user"

	"zerodemo/video/api/internal/svc"
	"zerodemo/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.ReqGetUser{Id: "1"})
	if err != nil {
		return nil, err
	}

	return &types.VideoResp{
		Id:   req.Id,
		Name: data.Name,
	}, nil
}
