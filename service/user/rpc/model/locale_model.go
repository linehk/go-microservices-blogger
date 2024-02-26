package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LocaleModel = (*customLocaleModel)(nil)

//go:generate mockgen -destination=./mock_locale_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/user/rpc/model github.com/linehk/go-microservices-blogger/service/user/rpc/model LocaleModel
type (
	// LocaleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLocaleModel.
	LocaleModel interface {
		localeModel
	}

	customLocaleModel struct {
		*defaultLocaleModel
	}
)

// NewLocaleModel returns a model for the database table.
func NewLocaleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LocaleModel {
	return &customLocaleModel{
		defaultLocaleModel: newLocaleModel(conn, c, opts...),
	}
}
