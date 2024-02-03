package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCenterModel = (*customUserCenterModel)(nil)

type (
	// UserCenterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCenterModel.
	UserCenterModel interface {
		userCenterModel
	}

	customUserCenterModel struct {
		*defaultUserCenterModel
	}
)

// NewUserCenterModel returns a model for the database table.
func NewUserCenterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserCenterModel {
	return &customUserCenterModel{
		defaultUserCenterModel: newUserCenterModel(conn, c, opts...),
	}
}
