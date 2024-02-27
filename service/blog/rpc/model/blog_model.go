package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogModel = (*customBlogModel)(nil)

//go:generate mockgen -destination=./mock_blog_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model BlogModel
type (
	// BlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogModel.
	BlogModel interface {
		blogModel
	}

	customBlogModel struct {
		*defaultBlogModel
	}
)

// NewBlogModel returns a model for the database table.
func NewBlogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BlogModel {
	return &customBlogModel{
		defaultBlogModel: newBlogModel(conn, c, opts...),
	}
}
