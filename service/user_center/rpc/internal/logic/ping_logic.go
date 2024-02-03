package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/user_center/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user_center/rpc/user_center"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user_center.Request) (*user_center.Response, error) {
	// todo: add your logic here and delete this line

	return &user_center.Response{}, nil
}
