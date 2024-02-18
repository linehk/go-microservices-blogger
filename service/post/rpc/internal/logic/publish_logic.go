package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *post.PublishReq) (*post.Post, error) {
	// todo: add your logic here and delete this line

	return &post.Post{}, nil
}
