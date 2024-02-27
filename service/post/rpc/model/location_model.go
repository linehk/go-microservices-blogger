package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LocationModel = (*customLocationModel)(nil)

//go:generate mockgen -destination=./mock_location_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model LocationModel
type (
	// LocationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLocationModel.
	LocationModel interface {
		locationModel
	}

	customLocationModel struct {
		*defaultLocationModel
	}
)

// NewLocationModel returns a model for the database table.
func NewLocationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LocationModel {
	return &customLocationModel{
		defaultLocationModel: newLocationModel(conn, c, opts...),
	}
}
