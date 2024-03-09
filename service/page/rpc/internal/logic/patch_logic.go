package logic

import (
	"context"
	"database/sql"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"

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
	pageReq := in.GetPage()
	pageModel := &model.Page{
		Uuid:      in.GetPageId(),
		BlogUuid:  sql.NullString{String: in.GetBlogId(), Valid: true},
		Status:    sql.NullString{String: pageReq.GetStatus(), Valid: true},
		Published: sql.NullTime{Time: pageReq.GetUpdated().AsTime(), Valid: true},
		Updated:   sql.NullTime{Time: pageReq.GetUpdated().AsTime(), Valid: true},
		Url:       sql.NullString{String: pageReq.GetUrl(), Valid: true},
		SelfLink:  sql.NullString{String: pageReq.GetSelfLink(), Valid: true},
		Title:     sql.NullString{String: pageReq.GetTitle(), Valid: true},
		Content:   sql.NullString{String: pageReq.GetContent(), Valid: true},
	}
	if err := l.svcCtx.PageModel.Update(l.ctx, pageModel); err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	return Get(l.ctx, l.svcCtx, l.Logger, pageModel)
}
