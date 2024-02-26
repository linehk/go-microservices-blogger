package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserModel = (*customAppUserModel)(nil)

//go:generate mockgen -destination=./mock_app_user_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/user/rpc/model github.com/linehk/go-microservices-blogger/service/user/rpc/model AppUserModel
type (
	// AppUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserModel.
	AppUserModel interface {
		appUserModel
	}

	customAppUserModel struct {
		*defaultAppUserModel
	}
)

// NewAppUserModel returns a model for the database table.
func NewAppUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AppUserModel {
	return &customAppUserModel{
		defaultAppUserModel: newAppUserModel(conn, c, opts...),
	}
}
