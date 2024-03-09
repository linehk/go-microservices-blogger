package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	commentmodel "github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

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

func (l *InsertLogic) Insert(in *post.InsertReq) (*post.Post, error) {
	postUuid := uuid.NewString()
	postReq := in.GetPost()
	postModel := &model.Post{
		Uuid:           postUuid,
		BlogUuid:       sql.NullString{String: in.GetBlogId(), Valid: true},
		Published:      sql.NullTime{Time: postReq.GetPublished().AsTime(), Valid: true},
		Updated:        sql.NullTime{Time: postReq.GetUpdated().AsTime(), Valid: true},
		Url:            postReq.GetUrl(),
		SelfLink:       sql.NullString{String: postReq.GetSelfLink(), Valid: true},
		Title:          sql.NullString{String: postReq.GetTitle(), Valid: true},
		TitleLink:      sql.NullString{String: postReq.GetTitleLink(), Valid: true},
		Content:        sql.NullString{String: postReq.GetContent(), Valid: true},
		CustomMetaData: sql.NullString{String: postReq.GetCustomMetaData(), Valid: true},
		Status:         sql.NullString{String: postReq.GetStatus(), Valid: true},
	}
	_, err := l.svcCtx.PostModel.Insert(l.ctx, postModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	for _, imageReq := range postReq.GetImages() {
		imageUuid := uuid.NewString()
		imageModel := &model.Image{
			Uuid:       imageUuid,
			PostUuid:   sql.NullString{String: postUuid, Valid: true},
			AuthorUuid: "",
			Url:        sql.NullString{String: imageReq.Url, Valid: true},
		}
		_, err := l.svcCtx.ImageModel.Insert(l.ctx, imageModel)
		if err != nil {
			l.Error(errcode.Msg(errcode.Database))
			return nil, errcode.Wrap(errcode.Database)
		}
	}

	authorUuid := uuid.NewString()
	authorModel := &model.Author{
		Uuid:        authorUuid,
		PostUuid:    postUuid,
		PageUuid:    "",
		CommentUuid: "",
		DisplayName: sql.NullString{String: postReq.GetAuthor().GetDisplayName(), Valid: true},
		Url:         sql.NullString{String: postReq.GetAuthor().GetUrl(), Valid: true},
	}
	_, err = l.svcCtx.AuthorModel.Insert(l.ctx, authorModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	authorImageUuid := uuid.NewString()
	authorImage := &model.Image{
		Uuid:       authorImageUuid,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid,
		Url:        sql.NullString{String: postReq.GetAuthor().GetImage().GetUrl(), Valid: true},
	}
	_, err = l.svcCtx.ImageModel.Insert(l.ctx, authorImage)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	for _, commentReq := range postReq.GetReplies().GetItems() {
		commentUuid := uuid.NewString()
		commentModel := &commentmodel.Comment{
			Uuid:      commentUuid,
			BlogUuid:  sql.NullString{String: in.GetBlogId(), Valid: true},
			PostUuid:  sql.NullString{String: postUuid, Valid: true},
			Status:    sql.NullString{String: commentReq.GetStatus(), Valid: true},
			Published: sql.NullTime{Time: commentReq.GetPublished().AsTime(), Valid: true},
			Updated:   sql.NullTime{Time: commentReq.GetUpdated().AsTime(), Valid: true},
			SelfLink:  sql.NullString{String: commentReq.GetSelfLink(), Valid: true},
			Content:   sql.NullString{String: commentReq.GetContent(), Valid: true},
		}
		_, err := l.svcCtx.CommentModel.Insert(l.ctx, commentModel)
		if err != nil {
			l.Error(errcode.Msg(errcode.Database))
			return nil, errcode.Wrap(errcode.Database)
		}

		commentAuthorUuid := uuid.NewString()
		commentAuthor := &model.Author{
			Uuid:        commentAuthorUuid,
			PostUuid:    "",
			PageUuid:    "",
			CommentUuid: commentUuid,
			DisplayName: sql.NullString{String: commentReq.GetAuthor().GetDisplayName(), Valid: true},
			Url:         sql.NullString{String: commentReq.GetAuthor().GetUrl(), Valid: true},
		}
		_, err = l.svcCtx.AuthorModel.Insert(l.ctx, commentAuthor)
		if err != nil {
			l.Error(errcode.Msg(errcode.Database))
			return nil, errcode.Wrap(errcode.Database)
		}

		commentAuthorImageUuid := uuid.NewString()
		commentAuthorImage := &model.Image{
			Uuid:       commentAuthorImageUuid,
			PostUuid:   sql.NullString{String: "", Valid: true},
			AuthorUuid: commentAuthorUuid,
			Url:        sql.NullString{String: commentReq.GetAuthor().GetImage().GetUrl(), Valid: true},
		}
		_, err = l.svcCtx.ImageModel.Insert(l.ctx, commentAuthorImage)
		if err != nil {
			l.Error(errcode.Msg(errcode.Database))
			return nil, errcode.Wrap(errcode.Database)
		}
	}

	for _, labelReq := range postReq.GetLabels() {
		labelUuid := uuid.NewString()
		labelModel := &model.Label{
			Uuid:       labelUuid,
			PostUuid:   sql.NullString{String: postUuid, Valid: true},
			LabelValue: sql.NullString{String: labelReq, Valid: true},
		}
		_, err = l.svcCtx.LabelModel.Insert(l.ctx, labelModel)
		if err != nil {
			l.Error(errcode.Msg(errcode.Database))
			return nil, errcode.Wrap(errcode.Database)
		}
	}

	locationUuid := uuid.NewString()
	locationModel := &model.Location{
		Uuid:     locationUuid,
		PostUuid: postUuid,
		Name:     sql.NullString{String: postReq.GetLocation().GetName(), Valid: true},
		Lat:      sql.NullFloat64{Float64: float64(postReq.GetLocation().GetLat()), Valid: true},
		Lng:      sql.NullFloat64{Float64: float64(postReq.GetLocation().GetLng()), Valid: true},
		Span:     sql.NullString{String: postReq.GetLocation().GetSpan(), Valid: true},
	}
	_, err = l.svcCtx.LocationModel.Insert(l.ctx, locationModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	newPostModel, err := l.svcCtx.PostModel.FindOneByBlogUuidAndPostUuid(l.ctx, in.GetBlogId(), postUuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, newPostModel)
}
