package logic

import (
	"context"
	"errors"

	"github.com/linehk/go-microservices-blogger/convert"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogUserInfosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlogUserInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogUserInfosLogic {
	return &GetBlogUserInfosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBlogUserInfosLogic) GetBlogUserInfos(in *blog.GetBlogUserInfosReq) (*blog.BlogUserInfos, error) {
	blogUserInfoModel, err := l.svcCtx.BlogUserInfoModel.FindOneByUserUuidAndBlogUuid(l.ctx, in.GetUserId(), in.GetBlogId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.BlogUserInfoNotExist))
		return nil, errcode.Wrap(errcode.BlogUserInfoNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	blogResp, err := NewGetLogic(l.ctx, l.svcCtx).Get(&blog.GetReq{
		BlogId: in.GetBlogId(),
	})
	if err != nil {
		l.Error(errcode.Msg(errcode.Service))
		return nil, errcode.Wrap(errcode.Service)
	}

	var blogUserInfos blog.BlogUserInfos
	blogUserInfos.Kind = "blogger#blogUserInfo"
	blogUserInfos.Blog = &blog.Blog{}
	blogUserInfos.BlogUserInfo = &blog.BlogUserInfo{}
	convert.Copy(&blogUserInfos.Blog, blogResp)
	convert.Copy(&blogUserInfos.BlogUserInfo, blogUserInfoModel)
	blogUserInfos.BlogUserInfo.Kind = "blogger#blogPerUserInfo"
	blogUserInfos.BlogUserInfo.UserId = blogUserInfoModel.UserUuid.String
	blogUserInfos.BlogUserInfo.BlogId = blogUserInfoModel.BlogUuid.String
	return &blogUserInfos, nil
}
