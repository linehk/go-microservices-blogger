package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveContentLogic {
	return &RemoveContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveContentLogic) RemoveContent(in *comment.RemoveContentReq) (*comment.Comment, error) {
	// todo: add your logic here and delete this line

	return &comment.Comment{}, nil
}
