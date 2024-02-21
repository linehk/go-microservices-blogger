package svc

import (
	"github.com/linehk/go-blogger/service/blog/rpc/internal/config"
	"github.com/linehk/go-blogger/service/blog/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	BlogModel         model.BlogModel
	BlogUserInfoModel model.BlogUserInfoModel
	PageViewsModel    model.PageViewsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Cache[0].Host,
			Type: redis.NodeType,
		}),
		BlogModel:         model.NewBlogModel(conn, c.Cache),
		BlogUserInfoModel: model.NewBlogUserInfoModel(conn, c.Cache),
		PageViewsModel:    model.NewPageViewsModel(conn, c.Cache),
	}
}
