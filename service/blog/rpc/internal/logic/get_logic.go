package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/blog/rpc/model"

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
		return nil, fmt.Errorf("BlogModel.FindOneByUuid NotFound err: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("BlogModel.FindOneByUuid err: %v", err)
	}

	var respBlog blog.Blog
	err = copier.Copy(&respBlog, blogModel)
	if err != nil {
		return nil, fmt.Errorf("copier.Copy err: %v", err)
	}
	respBlog.Kind = "blogger#blog"

	return &respBlog, nil
}
