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

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *comment.ListReq) (*comment.ListResp, error) {
	commentModelList, err := l.svcCtx.CommentModel.ListByBlogUuidAndPostUuid(l.ctx, in.GetBlogId(), in.GetPostId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.CommentNotExist))
		return nil, errcode.Wrap(errcode.CommentNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listResp comment.ListResp
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
