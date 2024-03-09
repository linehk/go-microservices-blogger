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

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *page.ListReq) (*page.ListResp, error) {
	pageModelList, err := l.svcCtx.PageModel.ListByBlogUuid(l.ctx, in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PageNotExist))
		return nil, errcode.Wrap(errcode.PageNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listResp page.ListResp
	listResp.Kind = "blogger#pageList"
	for _, pageModel := range pageModelList {
		pageResp, err := Get(l.ctx, l.svcCtx, l.Logger, pageModel)
		if err != nil {
			return nil, err
		}
		listResp.Items = append(listResp.Items, pageResp)
	}

	return &listResp, nil
}
