package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByBlogLogic {
	return &ListByBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByBlogLogic) ListByBlog(in *comment.ListByBlogReq) (*comment.ListByBlogResp, error) {
	commentModelList, err := l.svcCtx.CommentModel.ListByBlogUuid(l.ctx, in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.CommentNotExist))
		return nil, errcode.Wrap(errcode.CommentNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listResp comment.ListByBlogResp
	listResp.Kind = "blogger#commentList"
	listResp.PrevPageToken = ""
	listResp.NextPageToken = ""
	for _, commentModel := range commentModelList {
		commentResp, err := Get(l.ctx, l.svcCtx, l.Logger, commentModel)
		if err != nil {
			return nil, err
		}
		listResp.Items = append(listResp.Items, commentResp)
	}
	return &listResp, nil
}
