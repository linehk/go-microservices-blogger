package svc

import (
	"github.com/linehk/go-blogger/service/user_center/api/internal/config"
	"github.com/linehk/go-blogger/service/user_center/rpc/user_center_client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserCenter user_center_client.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserCenter: user_center_client.NewUserCenter(zrpc.MustNewClient(c.UserCenter)),
	}
}
