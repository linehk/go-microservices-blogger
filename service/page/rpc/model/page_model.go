package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PageModel = (*customPageModel)(nil)

//go:generate mockgen -destination=./mock_page_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/page/rpc/model github.com/linehk/go-microservices-blogger/service/page/rpc/model PageModel
type (
	// PageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPageModel.
	PageModel interface {
		pageModel
	}

	customPageModel struct {
		*defaultPageModel
	}
)

// NewPageModel returns a model for the database table.
func NewPageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PageModel {
	return &customPageModel{
		defaultPageModel: newPageModel(conn, c, opts...),
	}
}
