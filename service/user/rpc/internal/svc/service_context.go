package svc

import (
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blogservice"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	AppUserModel model.AppUserModel
	LocaleModel  model.LocaleModel
	BlogService  blogservice.BlogService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	redisConf := redis.RedisConf{
		Host: c.Cache[0].Host,
		Type: redis.NodeType,
	}
	
	return &ServiceContext{
		Config:       c,
		RedisClient:  redis.MustNewRedis(redisConf),
		AppUserModel: model.NewAppUserModel(conn, c.Cache),
		LocaleModel:  model.NewLocaleModel(conn, c.Cache),
		BlogService:  blogservice.NewBlogService(zrpc.MustNewClient(c.BlogConf)),
	}
}
