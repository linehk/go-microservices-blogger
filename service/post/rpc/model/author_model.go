package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthorModel = (*customAuthorModel)(nil)

//go:generate mockgen -destination=./mock_author_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model AuthorModel
type (
	// AuthorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthorModel.
	AuthorModel interface {
		authorModel
	}

	customAuthorModel struct {
		*defaultAuthorModel
	}
)

// NewAuthorModel returns a model for the database table.
func NewAuthorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthorModel {
	return &customAuthorModel{
		defaultAuthorModel: newAuthorModel(conn, c, opts...),
	}
}
