package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthorModel = (*customAuthorModel)(nil)

//go:generate mockgen -destination=./mock_author_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model AuthorModel
type (
	// AuthorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthorModel.
	AuthorModel interface {
		authorModel
		ListByPostUuid(ctx context.Context, postUuid string) ([]*Author, error)
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

var (
	cachePublicAuthorListPostUuidPrefix = "cache:public:author:list:postUuid:"
)

func (c *customAuthorModel) ListByPostUuid(ctx context.Context, postUuid string) ([]*Author, error) {
	publicAuthorListPostUuidKey := fmt.Sprintf("%s%v", cachePublicAuthorListPostUuidPrefix, postUuid)
	var resp []*Author
	err := c.QueryRowIndexCtx(ctx, &resp, publicAuthorListPostUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where post_uuid = $1", authorRows, c.table)
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
