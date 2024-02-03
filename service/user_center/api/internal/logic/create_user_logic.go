package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/user_center/api/internal/svc"
	"github.com/linehk/go-blogger/service/user_center/api/internal/types"
	"github.com/linehk/go-blogger/service/user_center/rpc/user_center"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) error {
	_, err := l.svcCtx.UserCenter.CreateUser(l.ctx, &user_center.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return err
	}
	return nil
}
