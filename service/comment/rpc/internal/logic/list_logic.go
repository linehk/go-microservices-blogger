package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-blogger/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *comment.ListReq) (*comment.ListResp, error) {
	// todo: add your logic here and delete this line

	return &comment.ListResp{}, nil
}
