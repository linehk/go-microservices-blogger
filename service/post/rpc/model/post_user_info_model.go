package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostUserInfoModel = (*customPostUserInfoModel)(nil)

type (
	// PostUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostUserInfoModel.
	PostUserInfoModel interface {
		postUserInfoModel
	}

	customPostUserInfoModel struct {
		*defaultPostUserInfoModel
	}
)

// NewPostUserInfoModel returns a model for the database table.
func NewPostUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostUserInfoModel {
	return &customPostUserInfoModel{
		defaultPostUserInfoModel: newPostUserInfoModel(conn, c, opts...),
	}
}
