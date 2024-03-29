package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
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

func (l *GetLogic) Get(in *post.GetReq) (*post.Post, error) {
	postModel, err := l.svcCtx.PostModel.FindOneByBlogUuidAndPostUuid(l.ctx, in.GetBlogId(), in.GetPostId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, postModel)
}

func Get(ctx context.Context, svcCtx *svc.ServiceContext, l logx.Logger, postModel *model.Post) (*post.Post, error) {
	var postResp post.Post
	convert.Copy(&postResp, postModel)
	postResp.Kind = "blogger#post"
	postResp.Id = postModel.Uuid
	postResp.Blog = &post.Blog{Id: postModel.BlogUuid.String}
	if postModel.Published.Valid {
		postResp.Published = timestamppb.New(postModel.Published.Time)
	}
	if postModel.Updated.Valid {
		postResp.Updated = timestamppb.New(postModel.Updated.Time)
	}

	imageModelList, err := svcCtx.ImageModel.ListByPostUuid(ctx, postModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.ImageNotExist))
		return nil, errcode.Wrap(errcode.ImageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	for _, imageModel := range imageModelList {
		if imageModel.Url.Valid {
			postResp.Images = append(postResp.Images, &post.Image{Url: imageModel.Url.String})
		}
	}

	authorModel, err := svcCtx.AuthorModel.FindOneByPostUuid(ctx, postModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.AuthorNotExist))
		return nil, errcode.Wrap(errcode.AuthorNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	postResp.Author = &post.Author{}
	convert.Copy(postResp.Author, authorModel)
	postResp.Author.Id = authorModel.Uuid
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
		postResp.Author.Image = &post.Image{Url: authorImageModel.Url.String}
	}

	listCommentReq := &comment.ListReq{
		BlogId: postModel.BlogUuid.String,
		PostId: postModel.Uuid,
	}
	listCommentResp, err := svcCtx.CommentService.List(ctx, listCommentReq)
	if err != nil {
		l.Error(errcode.Msg(errcode.Service))
		return nil, errcode.Wrap(errcode.Service)
	}

	postResp.Replies = &post.Reply{}
	postResp.Replies.TotalItems = int64(len(listCommentResp.GetItems()))
	for _, commentItem := range listCommentResp.GetItems() {
		var repliesItem post.Comment
		convert.Copy(&repliesItem, commentItem)
		postResp.Replies.Items = append(postResp.Replies.Items, &repliesItem)
	}

	labelModelList, err := svcCtx.LabelModel.ListByPostUuid(ctx, postModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.LabelNotExist))
		return nil, errcode.Wrap(errcode.LabelNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	for _, labelModel := range labelModelList {
		if labelModel.LabelValue.Valid {
			postResp.Labels = append(postResp.Labels, labelModel.LabelValue.String)
		}
	}

	locationModel, err := svcCtx.LocationModel.FindOneByPostUuid(ctx, postModel.Uuid)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.LocationNotExist))
		return nil, errcode.Wrap(errcode.LocationNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	postResp.Location = &post.Location{}
	convert.Copy(postResp.Location, locationModel)

	return &postResp, nil
}
