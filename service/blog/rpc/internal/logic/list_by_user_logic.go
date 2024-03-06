package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByUserLogic {
	return &ListByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByUserLogic) ListByUser(in *blog.ListByUserReq) (*blog.ListByUserResp, error) {
	blogModelList, err := l.svcCtx.BlogModel.ListByAppUserUuid(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.BlogNotExist))
		return nil, errcode.Wrap(errcode.BlogNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var listByUserResp blog.ListByUserResp
	listByUserResp.Kind = "blogger#blogList"

	for _, blogModel := range blogModelList {
		blogResp, err := Get(l.ctx, l.svcCtx, l.Logger, blogModel)
		if err != nil {
			return nil, err
		}
		listByUserResp.Items = append(listByUserResp.Items, blogResp)
	}

	return &listByUserResp, nil
}
