package logic

import (
	"context"

	"github.com/linehk/go-blogger/service/user_center/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user_center/rpc/user_center"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_center.CreateUserRequest) (*user_center.EmptyResponse, error) {
	// todo: add your logic here and delete this line

	return &user_center.EmptyResponse{}, nil
}
