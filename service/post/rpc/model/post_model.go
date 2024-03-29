package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

//go:generate mockgen -destination=./mock_post_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model PostModel
type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		FindOneByBlogUuidAndPostUuid(ctx context.Context, blogUuid, postUuid string) (*Post, error)
		ListByBlogUuid(ctx context.Context, blogUuid string) ([]*Post, error)
		SearchByTitle(ctx context.Context, blogUuid, title string) ([]*Post, error)
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn, c, opts...),
	}
}

var (
	cachePublicPostBlogUuidAndPostUuid = "cache:public:post:blogUuid:%s:postUuid:%s"
	cachePublicPostBlogUuidPrefix      = "cache:public:post:blogUuid:"
	cachePublicPostBlogUuidAndTitle    = "cache:public:post:blogUuid:%s:title:%s"
)

func (c *customPostModel) FindOneByBlogUuidAndPostUuid(ctx context.Context, blogUuid, postUuid string) (*Post, error) {
	publicPostBlogUuidAndPostUuidKey := fmt.Sprintf(cachePublicPostBlogUuidAndPostUuid, blogUuid, postUuid)
	var resp Post
	err := c.QueryRowIndexCtx(ctx, &resp, publicPostBlogUuidAndPostUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1 and blog_uuid = $2 limit 1", postRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid, postUuid); err != nil {
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

func (c *customPostModel) ListByBlogUuid(ctx context.Context, blogUuid string) ([]*Post, error) {
	publicPostBlogUuidKey := fmt.Sprintf("%s%v", cachePublicPostBlogUuidPrefix, blogUuid)
	var resp []*Post
	err := c.QueryRowIndexCtx(ctx, &resp, publicPostBlogUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1", postRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid); err != nil {
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

func (c *customPostModel) SearchByTitle(ctx context.Context, blogUuid, title string) ([]*Post, error) {
	publicPostBlogUuidAndTitleKey := fmt.Sprintf(cachePublicPostBlogUuidAndTitle, blogUuid, title)
	var resp []*Post
	err := c.QueryRowIndexCtx(ctx, &resp, publicPostBlogUuidAndTitleKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1 and title like $2", postRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid, "'"+title+"%'"); err != nil {
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
