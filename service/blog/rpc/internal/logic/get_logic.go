package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *blog.GetReq) (*blog.Blog, error) {
	blogModel, err := l.svcCtx.BlogModel.FindOneByUuid(l.ctx, in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		wrapErr := fmt.Errorf("BlogModel.FindOneByUuid NotFound err: %v", err)
		l.Error(wrapErr)
		return nil, wrapErr
	}
	if err != nil {
		wrapErr := fmt.Errorf("BlogModel.FindOneByUuid err: %v", err)
		l.Error(wrapErr)
		return nil, wrapErr
	}

	var blogResp blog.Blog
	err = copier.Copy(&blogResp, blogModel)
	if err != nil {
		wrapErr := fmt.Errorf("copier.Copy err: %v", err)
		l.Error(wrapErr)
		return nil, wrapErr
	}
	blogResp.Kind = "blogger#blog"

	return &blogResp, nil
}
