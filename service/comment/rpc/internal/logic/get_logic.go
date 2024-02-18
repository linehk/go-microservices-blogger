package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-blogger/service/comment/rpc/internal/svc"

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

func (l *GetLogic) Get(in *comment.GetReq) (*comment.Comment, error) {
	// todo: add your logic here and delete this line

	return &comment.Comment{}, nil
}
