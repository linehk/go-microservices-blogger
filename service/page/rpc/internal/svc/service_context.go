package svc

import (
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	postmodel "github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	PageModel   model.PageModel
	AuthorModel postmodel.AuthorModel
	ImageModel  postmodel.ImageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Cache[0].Host,
			Type: redis.NodeType,
		}),
		PageModel:   model.NewPageModel(conn, c.Cache),
		AuthorModel: postmodel.NewAuthorModel(conn, c.Cache),
		ImageModel:  postmodel.NewImageModel(conn, c.Cache),
	}
}
