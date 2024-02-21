package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchLogic {
	return &PatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchLogic) Patch(in *post.PatchReq) (*post.Post, error) {
	// todo: add your logic here and delete this line

	return &post.Post{}, nil
}
