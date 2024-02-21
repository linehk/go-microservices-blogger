package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByPathLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByPathLogic {
	return &GetByPathLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByPathLogic) GetByPath(in *post.GetByPathReq) (*post.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &post.EmptyResp{}, nil
}
