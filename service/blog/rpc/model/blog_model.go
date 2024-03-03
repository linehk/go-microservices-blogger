package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogModel = (*customBlogModel)(nil)

//go:generate mockgen -destination=./mock_blog_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model BlogModel
type (
	// BlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogModel.
	BlogModel interface {
		blogModel
		FindOneByUrl(ctx context.Context, url string) (*Blog, error)
	}

	customBlogModel struct {
		*defaultBlogModel
	}
)

// NewBlogModel returns a model for the database table.
func NewBlogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BlogModel {
	return &customBlogModel{
		defaultBlogModel: newBlogModel(conn, c, opts...),
	}
}

var (
	cachePublicBlogUrlPrefix = "cache:public:blog:url:"
)

func (c *customBlogModel) FindOneByUrl(ctx context.Context, url string) (*Blog, error) {
	publicBlogUrlKey := fmt.Sprintf("%s%v", cachePublicBlogUrlPrefix, url)
	var resp Blog
	err := c.QueryRowIndexCtx(ctx, &resp, publicBlogUrlKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where url = $1 limit 1", blogRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, url); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, c.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
