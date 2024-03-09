package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
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

func (l *GetLogic) Get(in *page.GetReq) (*page.Page, error) {
	pageModel, err := l.svcCtx.PageModel.FindOneByBlogUuidAndPageUuid(l.ctx, in.GetBlogId(), in.GetPageId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PageNotExist))
		return nil, errcode.Wrap(errcode.PageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, pageModel)
}

func Get(ctx context.Context, svcCtx *svc.ServiceContext, l logx.Logger, pageModel *model.Page) (*page.Page, error) {
	var pageResp page.Page
	convert.Copy(&pageResp, pageModel)
	pageResp.Kind = "blogger#page"
	pageResp.Id = pageModel.Uuid
	pageResp.Blog = &page.Blog{Id: pageModel.BlogUuid.String}
	if pageModel.Published.Valid {
		pageResp.Published = timestamppb.New(pageModel.Published.Time)
	}
	if pageModel.Updated.Valid {
		pageResp.Updated = timestamppb.New(pageModel.Updated.Time)
	}

	authorModel, err := svcCtx.AuthorModel.FindOneByPageUuid(ctx, pageModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.AuthorNotExist))
		return nil, errcode.Wrap(errcode.AuthorNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	pageResp.Author = &page.Author{}
	convert.Copy(pageResp.Author, authorModel)
	pageResp.Author.Id = authorModel.Uuid
	authorImageModel, err := svcCtx.ImageModel.FindOneByAuthorUuid(ctx, authorModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.ImageNotExist))
		return nil, errcode.Wrap(errcode.ImageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	if authorImageModel.Url.Valid {
		pageResp.Author.Image = &page.Image{Url: authorImageModel.Url.String}
	}

	return &pageResp, nil
}
