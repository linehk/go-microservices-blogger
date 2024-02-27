package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LabelModel = (*customLabelModel)(nil)

//go:generate mockgen -destination=./mock_label_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model LabelModel
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
