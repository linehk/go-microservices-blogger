package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LabelModel = (*customLabelModel)(nil)

type (
	// LabelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLabelModel.
	LabelModel interface {
		labelModel
	}

	customLabelModel struct {
		*defaultLabelModel
	}
)

// NewLabelModel returns a model for the database table.
func NewLabelModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LabelModel {
	return &customLabelModel{
		defaultLabelModel: newLabelModel(conn, c, opts...),
	}
}
