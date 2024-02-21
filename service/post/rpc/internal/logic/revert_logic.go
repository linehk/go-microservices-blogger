package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevertLogic {
	return &RevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RevertLogic) Revert(in *post.RevertReq) (*post.Post, error) {
	// todo: add your logic here and delete this line

	return &post.Post{}, nil
}
