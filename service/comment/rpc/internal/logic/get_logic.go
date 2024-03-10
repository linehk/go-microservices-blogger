package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (l *GetLogic) Get(in *comment.GetReq) (*comment.Comment, error) {
	commentModel, err := l.svcCtx.CommentModel.FindOneByUuid(l.ctx, in.GetCommentId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.CommentNotExist))
		return nil, errcode.Wrap(errcode.CommentNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, commentModel)
}

func Get(ctx context.Context, svcCtx *svc.ServiceContext, l logx.Logger, commentModel *model.Comment) (*comment.Comment, error) {
	var commentResp comment.Comment
	convert.Copy(&commentResp, commentModel)
	commentResp.Kind = "blogger#comment"
	commentResp.Id = commentModel.Uuid
	commentResp.Post = &comment.Post{Id: commentModel.PostUuid.String}
	commentResp.Blog = &comment.Blog{Id: commentModel.BlogUuid.String}
	if commentModel.Published.Valid {
		commentResp.Published = timestamppb.New(commentModel.Published.Time)
	}
	if commentModel.Updated.Valid {
		commentResp.Updated = timestamppb.New(commentModel.Updated.Time)
	}

	authorModel, err := svcCtx.AuthorModel.FindOneByCommentUuid(ctx, commentModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.AuthorNotExist))
		return nil, errcode.Wrap(errcode.AuthorNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	commentResp.Author = &comment.Author{}
	convert.Copy(commentResp.Author, authorModel)
	commentResp.Author.Id = authorModel.Uuid
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
		commentResp.Author.Image = &comment.Image{Url: authorImageModel.Url.String}
	}

	return &commentResp, nil
}
