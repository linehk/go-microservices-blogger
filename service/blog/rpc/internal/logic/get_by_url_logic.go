package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"
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
	blogModel, err := l.svcCtx.BlogModel.FindOneByUrl(l.ctx, in.GetUrl())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.BlogNotExist))
		return nil, errcode.Wrap(errcode.BlogNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	return Get(l.ctx, l.svcCtx, l.Logger, blogModel)
}
