package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostUserInfosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostUserInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostUserInfosLogic {
	return &GetPostUserInfosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostUserInfosLogic) GetPostUserInfos(in *post.GetPostUserInfosReq) (*post.PostUserInfos, error) {
	postUserInfosModel, err := l.svcCtx.PostUserInfoModel.FindOneByUserUuidAndBlogUuidAndPostUuid(l.ctx, in.GetUserId(), in.GetBlogId(), in.GetPostId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostUserInfosNotExist))
		return nil, errcode.Wrap(errcode.PostUserInfosNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return GetPostUserInfos(l.ctx, l.svcCtx, l, postUserInfosModel)
}

func GetPostUserInfos(ctx context.Context, svcCtx *svc.ServiceContext, l logx.Logger, postUserInfosModel *model.PostUserInfo) (*post.PostUserInfos, error) {
	var postUserInfosResp post.PostUserInfos
	postUserInfosResp.Kind = "blogger#postUserInfo"
	postUserInfosResp.PostUserInfo = &post.PostUserInfo{}
	convert.Copy(postUserInfosResp.PostUserInfo, postUserInfosModel)
	postUserInfosResp.PostUserInfo.Kind = "blogger#postPerUserInfo"
	postUserInfosResp.PostUserInfo.UserId = postUserInfosModel.UserUuid.String
	postUserInfosResp.PostUserInfo.BlogId = postUserInfosModel.BlogUuid.String
	postUserInfosResp.PostUserInfo.PostId = postUserInfosModel.PostUuid.String

	postModel, err := svcCtx.PostModel.FindOneByBlogUuidAndPostUuid(ctx, postUserInfosModel.BlogUuid.String, postUserInfosModel.PostUuid.String)
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	postResp, err := Get(ctx, svcCtx, l, postModel)
	if err != nil {
		return nil, err
	}
	postUserInfosResp.Post = postResp

	return &postUserInfosResp, nil
}
