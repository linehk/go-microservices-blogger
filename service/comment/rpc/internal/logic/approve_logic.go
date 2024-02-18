package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-blogger/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApproveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveLogic {
	return &ApproveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApproveLogic) Approve(in *comment.ApproveReq) (*comment.Comment, error) {
	// todo: add your logic here and delete this line

	return &comment.Comment{}, nil
}
