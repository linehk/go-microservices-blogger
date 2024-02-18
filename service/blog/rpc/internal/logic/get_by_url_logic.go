package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-blogger/service/blog/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByUrlLogic {
	return &GetByUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByUrlLogic) GetByUrl(in *blog.GetByUrlReq) (*blog.Blog, error) {
	// todo: add your logic here and delete this line

	return &blog.Blog{}, nil
}
