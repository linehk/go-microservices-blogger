package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogUserInfoModel = (*customBlogUserInfoModel)(nil)

//go:generate mockgen -destination=./mock_blog_user_info_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model BlogUserInfoModel
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
