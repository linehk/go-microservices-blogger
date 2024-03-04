package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LabelModel = (*customLabelModel)(nil)

//go:generate mockgen -destination=./mock_label_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model LabelModel
type (
	// LabelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLabelModel.
	LabelModel interface {
		labelModel
		ListByPostUuid(ctx context.Context, postUuid string) ([]*Label, error)
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

var (
	cachePublicLabelListPostUuidPrefix = "cache:public:label:postUuid:"
)

func (c *customLabelModel) ListByPostUuid(ctx context.Context, postUuid string) ([]*Label, error) {
	publicLabelListPostUuidKey := fmt.Sprintf("%s%v", cachePublicLabelListPostUuidPrefix, postUuid)
	var resp []*Label
	err := c.QueryRowIndexCtx(ctx, &resp, publicLabelListPostUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where post_uuid = $1", labelRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, postUuid); err != nil {
			return nil, err
		}
		return resp[0].Id, nil
	}, c.queryPrimary)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
