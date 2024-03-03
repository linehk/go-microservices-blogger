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

type GetPageViewsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageViewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageViewsLogic {
	return &GetPageViewsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageViewsLogic) GetPageViews(in *blog.GetPageViewsReq) (*blog.PageViews, error) {
	pageViewsModel, err := l.svcCtx.PageViewsModel.FindOneByBlogUuid(l.ctx, in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PageViewNotExist))
		return nil, errcode.Wrap(errcode.PageViewNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var pageViewsResp blog.PageViews
	pageViewsResp.Kind = "blogger#page_views"
	pageViewsResp.BlogId = pageViewsModel.BlogUuid
	if pageViewsModel.Count.Valid {
		pageViewsResp.Counts = append(pageViewsResp.Counts, &blog.Count{
			TimeRange: "all",
			Count:     uint64(pageViewsModel.Count.Int64),
		})
	}

	return &pageViewsResp, nil
}
