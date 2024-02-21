package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
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
	// todo: add your logic here and delete this line

	return &post.ListPostUserInfosResp{}, nil
}
