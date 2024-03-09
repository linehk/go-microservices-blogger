package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"

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

func (l *DeleteLogic) Delete(in *page.DeleteReq) (*page.EmptyResp, error) {
	pageModel, err := l.svcCtx.PageModel.FindOneByBlogUuidAndPageUuid(l.ctx, in.GetBlogId(), in.GetPageId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PageNotExist))
		return nil, errcode.Wrap(errcode.PageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	if err := l.svcCtx.PageModel.Delete(l.ctx, pageModel.Id); err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return &page.EmptyResp{}, nil
}
