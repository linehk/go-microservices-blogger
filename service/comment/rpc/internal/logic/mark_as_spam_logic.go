package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-blogger/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkAsSpamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkAsSpamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkAsSpamLogic {
	return &MarkAsSpamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkAsSpamLogic) MarkAsSpam(in *comment.MarkAsSpamReq) (*comment.Comment, error) {
	// todo: add your logic here and delete this line

	return &comment.Comment{}, nil
}
