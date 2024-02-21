package logic

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *page.UpdateReq) (*page.Page, error) {
	// todo: add your logic here and delete this line

	return &page.Page{}, nil
}
