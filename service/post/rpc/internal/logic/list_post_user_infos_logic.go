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

type ListPostUserInfosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPostUserInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostUserInfosLogic {
	return &ListPostUserInfosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPostUserInfosLogic) ListPostUserInfos(in *post.ListPostUserInfosReq) (*post.ListPostUserInfosResp, error) {
	postUserInfosModelList, err := l.svcCtx.PostUserInfoModel.ListByUserUuidAndBlogUuid(l.ctx, in.GetUserId(), in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostUserInfosNotExist))
		return nil, errcode.Wrap(errcode.PostUserInfosNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listPostUserInfosResp post.ListPostUserInfosResp
	listPostUserInfosResp.Kind = "blogger#postUserInfosList"
	listPostUserInfosResp.NextPageToken = ""
	for _, postUserInfosModel := range postUserInfosModelList {
		postUserInfosResp, err := GetPostUserInfos(l.ctx, l.svcCtx, l.Logger, postUserInfosModel)
		if err != nil {
			return nil, err
		}
		listPostUserInfosResp.Items = append(listPostUserInfosResp.Items, postUserInfosResp)
	}

	return &listPostUserInfosResp, nil
}
