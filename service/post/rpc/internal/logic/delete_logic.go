package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *post.DeleteReq) (*post.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &post.EmptyResp{}, nil
}
