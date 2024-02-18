package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-blogger/service/blog/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogUserInfosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlogUserInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogUserInfosLogic {
	return &GetBlogUserInfosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBlogUserInfosLogic) GetBlogUserInfos(in *blog.BlogUserInfosReq) (*blog.BlogUserInfos, error) {
	// todo: add your logic here and delete this line

	return &blog.BlogUserInfos{}, nil
}
