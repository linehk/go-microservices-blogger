package svc

import (
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Cache[0].Host,
			Type: redis.NodeType,
		}),
		CommentModel: model.NewCommentModel(conn, c.Cache),
	}
}
