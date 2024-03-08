package logic

import (
	"context"
	"database/sql"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
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
	postReq := in.GetPost()
	postModel := &model.Post{
		Uuid:           in.GetPostId(),
		BlogUuid:       sql.NullString{String: in.GetBlogId(), Valid: true},
		Published:      sql.NullTime{Time: postReq.GetPublished().AsTime(), Valid: true},
		Updated:        sql.NullTime{Time: postReq.GetUpdated().AsTime(), Valid: true},
		Url:            postReq.GetUrl(),
		SelfLink:       sql.NullString{String: postReq.GetSelfLink(), Valid: true},
		Title:          sql.NullString{String: postReq.GetTitle(), Valid: true},
		TitleLink:      sql.NullString{String: postReq.GetTitleLink(), Valid: true},
		Content:        sql.NullString{String: postReq.GetContent(), Valid: true},
		CustomMetaData: sql.NullString{String: postReq.GetCustomMetaData(), Valid: true},
		Status:         sql.NullString{String: postReq.GetStatus(), Valid: true},
	}
	if err := l.svcCtx.PostModel.Update(l.ctx, postModel); err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	
	return Get(l.ctx, l.svcCtx, l.Logger, postModel)
}
