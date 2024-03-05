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

func (l *DeleteLogic) Delete(in *post.DeleteReq) (*post.EmptyResp, error) {
	postModel, err := l.svcCtx.PostModel.FindOneByUuid(l.ctx, in.GetPostId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	if postModel.BlogUuid != in.GetBlogId() {
		l.Error(errcode.Msg(errcode.PostNotBelongToBlog))
		return nil, errcode.Wrap(errcode.PostNotBelongToBlog)
	}
	if err := l.svcCtx.PostModel.Delete(l.ctx, postModel.Id); err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return &post.EmptyResp{}, nil
}
