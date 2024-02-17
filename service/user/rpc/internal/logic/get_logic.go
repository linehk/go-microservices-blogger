package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *user.GetReq) (*user.User, error) {
	// todo: add your logic here and delete this line

	return &user.User{}, nil
}
