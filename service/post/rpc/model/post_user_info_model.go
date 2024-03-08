package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostUserInfoModel = (*customPostUserInfoModel)(nil)

//go:generate mockgen -destination=./mock_post_user_info_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model PostUserInfoModel
type (
	// PostUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostUserInfoModel.
	PostUserInfoModel interface {
		postUserInfoModel
		FindOneByUserUuidAndBlogUuidAndPostUuid(ctx context.Context, userUuid, blogUuid, postUuid string) (*PostUserInfo, error)
		ListByUserUuidAndBlogUuid(ctx context.Context, userUuid, blogUuid string) ([]*PostUserInfo, error)
	}

	customPostUserInfoModel struct {
		*defaultPostUserInfoModel
	}
)

// NewPostUserInfoModel returns a model for the database table.
func NewPostUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostUserInfoModel {
	return &customPostUserInfoModel{
		defaultPostUserInfoModel: newPostUserInfoModel(conn, c, opts...),
	}
}

var (
	cachePublicPostUserInfoUserUuidBlogUuidPostUuid = "cache:public:postUserInfo:userUuid:%s:blogUuid:%s:postUuid:%s"
	cachePublicPostUserInfoUserUuidBlogUuid         = "cache:public:postUserInfo:userUuid:%s:blogUuid:%s"
)

func (c *customPostUserInfoModel) FindOneByUserUuidAndBlogUuidAndPostUuid(ctx context.Context, userUuid, blogUuid, postUuid string) (*PostUserInfo, error) {
	publicPostUserInfoBlogUuidKey := fmt.Sprintf(cachePublicPostUserInfoUserUuidBlogUuidPostUuid, userUuid, blogUuid, postUuid)
	var resp PostUserInfo
	err := c.QueryRowIndexCtx(ctx, &resp, publicPostUserInfoBlogUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where user_uuid = $1 and blog_uuid = $2 and post_uuid = $3 limit 1", postUserInfoRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userUuid, blogUuid, postUuid); err != nil {
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

func (c *customPostUserInfoModel) ListByUserUuidAndBlogUuid(ctx context.Context, userUuid, blogUuid string) ([]*PostUserInfo, error) {
	publicPostUserInfoBlogUuidKey := fmt.Sprintf(cachePublicPostUserInfoUserUuidBlogUuid, userUuid, blogUuid)
	var resp []*PostUserInfo
	err := c.QueryRowIndexCtx(ctx, &resp, publicPostUserInfoBlogUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where user_uuid = $1 and blog_uuid = $2", postUserInfoRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userUuid, blogUuid); err != nil {
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
