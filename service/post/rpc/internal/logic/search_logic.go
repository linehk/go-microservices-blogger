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

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *post.SearchReq) (*post.SearchResp, error) {
	postModelList, err := l.svcCtx.PostModel.SearchByTitle(l.ctx, in.GetBlogId(), in.GetQ())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	
	var searchResp post.SearchResp
	searchResp.Kind = "blogger#postList"
	searchResp.NextPageToken = ""
	for _, postModel := range postModelList {
		postResp, err := Get(l.ctx, l.svcCtx, l.Logger, postModel)
		if err != nil {
			return nil, err
		}
		searchResp.Items = append(searchResp.Items, postResp)
	}

	return &searchResp, nil
}
