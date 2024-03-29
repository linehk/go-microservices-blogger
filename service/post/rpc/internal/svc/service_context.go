package svc

import (
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/commentservice"
	commentmodel "github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	AuthorModel       model.AuthorModel
	ImageModel        model.ImageModel
	LabelModel        model.LabelModel
	LocationModel     model.LocationModel
	PostModel         model.PostModel
	PostUserInfoModel model.PostUserInfoModel
	CommentService    commentservice.CommentService
	CommentModel      commentmodel.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Cache[0].Host,
			Type: redis.NodeType,
		}),
		AuthorModel:       model.NewAuthorModel(conn, c.Cache),
		ImageModel:        model.NewImageModel(conn, c.Cache),
		LabelModel:        model.NewLabelModel(conn, c.Cache),
		LocationModel:     model.NewLocationModel(conn, c.Cache),
		PostModel:         model.NewPostModel(conn, c.Cache),
		PostUserInfoModel: model.NewPostUserInfoModel(conn, c.Cache),
		CommentService:    commentservice.NewCommentService(zrpc.MustNewClient(c.CommentConf)),
		CommentModel:      commentmodel.NewCommentModel(conn, c.Cache),
	}
}
