package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
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
	// todo: add your logic here and delete this line

	return &post.PostUserInfos{}, nil
}
