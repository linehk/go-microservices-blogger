package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/page/rpc/page"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *page.InsertReq) (*page.Page, error) {
	// todo: add your logic here and delete this line

	return &page.Page{}, nil
}
