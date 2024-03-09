package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PageModel = (*customPageModel)(nil)

//go:generate mockgen -destination=./mock_page_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/page/rpc/model github.com/linehk/go-microservices-blogger/service/page/rpc/model PageModel
type (
	// PageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPageModel.
	PageModel interface {
		pageModel
		FindOneByBlogUuidAndPageUuid(ctx context.Context, blogUuid, pageUuid string) (*Page, error)
	}

	customPageModel struct {
		*defaultPageModel
	}
)

// NewPageModel returns a model for the database table.
func NewPageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PageModel {
	return &customPageModel{
		defaultPageModel: newPageModel(conn, c, opts...),
	}
}

var (
	cachePublicBlogUuidAndPageUuid = "cache:public:page:blogUuid:%s:pageUuid:%s"
)

func (c *customPageModel) FindOneByBlogUuidAndPageUuid(ctx context.Context, blogUuid, pageUuid string) (*Page, error) {
	publicPageBlogUuidAndPageUuidKey := fmt.Sprintf(cachePublicBlogUuidAndPageUuid, blogUuid, pageUuid)
	var resp Page
	err := c.QueryRowIndexCtx(ctx, &resp, publicPageBlogUuidAndPageUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1 and page_uuid = $2 limit 1", pageRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid, pageUuid); err != nil {
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
