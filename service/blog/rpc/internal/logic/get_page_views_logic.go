package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageViewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageViewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageViewsLogic {
	return &GetPageViewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageViewsLogic) GetPageViews(in *blog.GetPageViewsReq) (*blog.PageViews, error) {
	// todo: add your logic here and delete this line

	return &blog.PageViews{}, nil
}
