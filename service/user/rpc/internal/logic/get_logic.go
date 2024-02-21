package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *user.GetReq) (*user.User, error) {
	appUserModel, err := l.svcCtx.AppUserModel.FindOneByUuid(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		return nil, fmt.Errorf("AppUserModel.FindOneByUuid NotFound err: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("AppUserModel.FindOneByUuid err: %v", err)
	}

	var userResp user.User
	err = copier.Copy(&userResp, appUserModel)
	if err != nil {
		return nil, fmt.Errorf("copier.Copy err: %v", err)
	}
	userResp.Id = appUserModel.Uuid
	userResp.Kind = "blogger#user"

	localeModel, err := l.svcCtx.LocaleModel.FindOneByAppUserUuid(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		return nil, fmt.Errorf("LocaleModel.FindOneByAppUserUuid NotFound err: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("LocaleModel.FindOneByAppUserUuid err: %v", err)
	}

	err = copier.Copy(&userResp.Locale, localeModel)
	if err != nil {
		return nil, fmt.Errorf("copier.Copy err: %v", err)
	}

	listByUserReq := &blog.ListByUserReq{
		UserId: in.GetUserId(),
	}
	listByUserResp, err := l.svcCtx.BlogService.ListByUser(l.ctx, listByUserReq)
	if err != nil {
		return nil, fmt.Errorf("BlogService.ListByUser err: %v", err)
	}
	for i, blogItem := range listByUserResp.GetItems() {
		userResp.Blogs[i].SelfLink = blogItem.GetSelfLink()
	}

	return &userResp, nil
}
