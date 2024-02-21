package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user/rpc/model"
	"github.com/linehk/go-blogger/service/user/rpc/user"

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
	userResp.Kind = "blogger#user"

	return &userResp, nil
}
