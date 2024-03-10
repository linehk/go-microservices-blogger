package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveContentLogic {
	return &RemoveContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveContentLogic) RemoveContent(in *comment.RemoveContentReq) (*comment.Comment, error) {
	commentModel, err := l.svcCtx.CommentModel.FindOneByUuid(l.ctx, in.GetCommentId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.CommentNotExist))
		return nil, errcode.Wrap(errcode.CommentNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	commentModel.Content.Valid = true
	commentModel.Content.String = ""

	if err := l.svcCtx.CommentModel.Update(l.ctx, commentModel); err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}
	return Get(l.ctx, l.svcCtx, l.Logger, commentModel)
}
