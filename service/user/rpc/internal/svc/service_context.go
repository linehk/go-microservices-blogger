package svc

import (
	"github.com/linehk/go-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-blogger/service/user/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	AppUserModel model.AppUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Cache[0].Host,
			Type: redis.NodeType,
		}),
		AppUserModel: model.NewAppUserModel(conn, c.Cache),
	}
}
