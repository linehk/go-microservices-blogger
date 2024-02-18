package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/page/rpc/page"

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

func (l *PatchLogic) Patch(in *page.PatchReq) (*page.Page, error) {
	// todo: add your logic here and delete this line

	return &page.Page{}, nil
}
