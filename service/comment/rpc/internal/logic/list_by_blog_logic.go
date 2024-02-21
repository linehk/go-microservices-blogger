package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByBlogLogic {
	return &ListByBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByBlogLogic) ListByBlog(in *comment.ListByBlogReq) (*comment.ListByBlogResp, error) {
	// todo: add your logic here and delete this line

	return &comment.ListByBlogResp{}, nil
}
