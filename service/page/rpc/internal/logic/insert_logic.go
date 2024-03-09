package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
	postmodel "github.com/linehk/go-microservices-blogger/service/post/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *page.InsertReq) (*page.Page, error) {
	pageUuid := uuid.NewString()
	pageReq := in.GetPage()
	pageModel := &model.Page{
		Uuid:      pageUuid,
		BlogUuid:  sql.NullString{String: in.GetBlogId(), Valid: true},
		Status:    sql.NullString{String: pageReq.GetStatus(), Valid: true},
		Published: sql.NullTime{Time: pageReq.GetUpdated().AsTime(), Valid: true},
		Updated:   sql.NullTime{Time: pageReq.GetUpdated().AsTime(), Valid: true},
		Url:       sql.NullString{String: pageReq.GetUrl(), Valid: true},
		SelfLink:  sql.NullString{String: pageReq.GetSelfLink(), Valid: true},
		Title:     sql.NullString{String: pageReq.GetTitle(), Valid: true},
		Content:   sql.NullString{String: pageReq.GetContent(), Valid: true},
	}
	_, err := l.svcCtx.PageModel.Insert(l.ctx, pageModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	authorUuid := uuid.NewString()
	authorModel := &postmodel.Author{
		Uuid:        authorUuid,
		PostUuid:    "",
		PageUuid:    pageUuid,
		CommentUuid: "",
		DisplayName: sql.NullString{String: pageReq.GetAuthor().GetDisplayName(), Valid: true},
		Url:         sql.NullString{String: pageReq.GetAuthor().GetUrl(), Valid: true},
	}
	_, err = l.svcCtx.AuthorModel.Insert(l.ctx, authorModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	authorImageUuid := uuid.NewString()
	authorImage := &postmodel.Image{
		Uuid:       authorImageUuid,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid,
		Url:        sql.NullString{String: pageReq.GetAuthor().GetImage().GetUrl(), Valid: true},
	}
	_, err = l.svcCtx.ImageModel.Insert(l.ctx, authorImage)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	newPageModel, err := l.svcCtx.PageModel.FindOneByBlogUuidAndPageUuid(l.ctx, in.GetBlogId(), pageUuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PageNotExist))
		return nil, errcode.Wrap(errcode.PageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, newPageModel)
}
