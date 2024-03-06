package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"google.golang.org/protobuf/types/known/timestamppb"

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
		l.Error(errcode.Msg(errcode.BlogNotExist))
		return nil, errcode.Wrap(errcode.BlogNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, blogModel)
}

func Get(ctx context.Context, svcCtx *svc.ServiceContext, l logx.Logger, blogModel *model.Blog) (*blog.Blog, error) {
	var blogResp blog.Blog
	convert.Copy(&blogResp, blogModel)
	blogResp.Kind = "blogger#blog"
	blogResp.Id = blogModel.Uuid
	if blogModel.Published.Valid {
		blogResp.Published = timestamppb.New(blogModel.Published.Time)
	}
	if blogModel.Updated.Valid {
		blogResp.Updated = timestamppb.New(blogModel.Updated.Time)
	}

	listPostReq := &post.ListReq{
		BlogId: blogModel.Uuid,
	}
	listPostResp, err := svcCtx.PostService.List(ctx, listPostReq)
	if err != nil {
		l.Error(errcode.Msg(errcode.Service))
		return nil, errcode.Wrap(errcode.Service)
	}
	postTotalItems := strconv.Itoa(len(listPostResp.GetItems()))
	for _, postItem := range listPostResp.GetItems() {
		blogResp.Posts = append(blogResp.Posts, &blog.Posts{TotalItems: postTotalItems, SelfLink: postItem.GetSelfLink()})
	}

	listPageReq := &page.ListReq{
		BlogId: blogModel.Uuid,
	}
	listPageResp, err := svcCtx.PageService.List(ctx, listPageReq)
	if err != nil {
		l.Error(errcode.Msg(errcode.Service))
		return nil, errcode.Wrap(errcode.Service)
	}
	pageTotalItems := strconv.Itoa(len(listPageResp.GetItems()))
	for _, pageItem := range listPageResp.GetItems() {
		blogResp.Pages = append(blogResp.Pages, &blog.Pages{TotalItems: pageTotalItems, SelfLink: pageItem.GetSelfLink()})
	}
	return &blogResp, nil
}
