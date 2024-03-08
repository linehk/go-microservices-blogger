package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

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

func (l *ListLogic) List(in *post.ListReq) (*post.ListResp, error) {
	postModelList, err := l.svcCtx.PostModel.ListByBlogUuid(l.ctx, in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listResp post.ListResp
	listResp.Kind = "blogger#postList"
	listResp.NextPageToken = ""
	for _, postModel := range postModelList {
		postResp, err := Get(l.ctx, l.svcCtx, l.Logger, postModel)
		if err != nil {
			return nil, err
		}
		listResp.Items = append(listResp.Items, postResp)
	}
	
	return &listResp, nil
}
