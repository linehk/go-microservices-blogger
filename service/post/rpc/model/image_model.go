package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ImageModel = (*customImageModel)(nil)

//go:generate mockgen -destination=./mock_image_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model ImageModel
type (
	// ImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImageModel.
	ImageModel interface {
		imageModel
	}

	customImageModel struct {
		*defaultImageModel
	}
)

// NewImageModel returns a model for the database table.
func NewImageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ImageModel {
	return &customImageModel{
		defaultImageModel: newImageModel(conn, c, opts...),
	}
}
