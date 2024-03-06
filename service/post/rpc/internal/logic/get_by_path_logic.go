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

type GetByPathLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByPathLogic {
	return &GetByPathLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByPathLogic) GetByPath(in *post.GetByPathReq) (*post.Post, error) {
	postModel, err := l.svcCtx.PostModel.FindOneByUrl(l.ctx, in.GetPath())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.PostNotExist))
		return nil, errcode.Wrap(errcode.PostNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	return Get(l.ctx, l.svcCtx, l.Logger, postModel)
}
