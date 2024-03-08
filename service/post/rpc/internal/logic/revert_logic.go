package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevertLogic {
	return &RevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RevertLogic) Revert(in *post.RevertReq) (*post.Post, error) {
	postModel, err := l.svcCtx.PostModel.FindOneByBlogUuidAndPostUuid(l.ctx, in.GetBlogId(), in.GetPostId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	postModel.Published.Valid = false

	if err := l.svcCtx.PostModel.Update(l.ctx, postModel); err != nil {
		return nil, err
	}

	return Get(l.ctx, l.svcCtx, l.Logger, postModel)
}
