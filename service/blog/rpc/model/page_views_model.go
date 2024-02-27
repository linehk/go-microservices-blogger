package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PageViewsModel = (*customPageViewsModel)(nil)

//go:generate mockgen -destination=./mock_page_views_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model PageViewsModel
type (
	// PageViewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPageViewsModel.
	PageViewsModel interface {
		pageViewsModel
	}

	customPageViewsModel struct {
		*defaultPageViewsModel
	}
)

// NewPageViewsModel returns a model for the database table.
func NewPageViewsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PageViewsModel {
	return &customPageViewsModel{
		defaultPageViewsModel: newPageViewsModel(conn, c, opts...),
	}
}
