package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogUserInfoModel = (*customBlogUserInfoModel)(nil)

type (
	// BlogUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogUserInfoModel.
	BlogUserInfoModel interface {
		blogUserInfoModel
	}

	customBlogUserInfoModel struct {
		*defaultBlogUserInfoModel
	}
)

// NewBlogUserInfoModel returns a model for the database table.
func NewBlogUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BlogUserInfoModel {
	return &customBlogUserInfoModel{
		defaultBlogUserInfoModel: newBlogUserInfoModel(conn, c, opts...),
	}
}
