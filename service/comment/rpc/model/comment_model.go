package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

//go:generate mockgen -destination=./mock_comment_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/comment/rpc/model github.com/linehk/go-microservices-blogger/service/comment/rpc/model CommentModel
type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		ListByBlogUuidAndPostUuid(ctx context.Context, blogUuid, postUuid string) ([]*Comment, error)
		ListByBlogUuid(ctx context.Context, blogUuid string) ([]*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c, opts...),
	}
}

var (
	cachePublicCommentBlogUuidAndPostUuid = "cache:public:comment:blogUuid:%s:postUuid:%s"
	cachePublicCommentBlogUuid            = "cache:public:comment:blogUuid:%s"
)

func (c *customCommentModel) ListByBlogUuidAndPostUuid(ctx context.Context, blogUuid, postUuid string) ([]*Comment, error) {
	publicCommentBlogUuidAndPostUuidKey := fmt.Sprintf(cachePublicCommentBlogUuidAndPostUuid, blogUuid, postUuid)
	var resp []*Comment
	err := c.QueryRowIndexCtx(ctx, &resp, publicCommentBlogUuidAndPostUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1 and post_uuid = $2", commentRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid, postUuid); err != nil {
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

func (c *customCommentModel) ListByBlogUuid(ctx context.Context, blogUuid string) ([]*Comment, error) {
	publicCommentBlogUuidKey := fmt.Sprintf(cachePublicCommentBlogUuid, blogUuid)
	var resp []*Comment
	err := c.QueryRowIndexCtx(ctx, &resp, publicCommentBlogUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1", commentRows, c.table)
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
