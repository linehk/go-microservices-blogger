package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogUserInfoModel = (*customBlogUserInfoModel)(nil)

//go:generate mockgen -destination=./mock_blog_user_info_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model BlogUserInfoModel
type (
	// BlogUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogUserInfoModel.
	BlogUserInfoModel interface {
		blogUserInfoModel
		FindOneByUserUuidAndBlogUuid(ctx context.Context, userUuid string, blogUuid string) (*BlogUserInfo, error)
	}

	customBlogUserInfoModel struct {
		*defaultBlogUserInfoModel
	}
)

// NewBlogUserInfoModel returns a model for the database table.
func NewBlogUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BlogUserInfoModel {
	return &customBlogUserInfoModel{
		defaultBlogUserInfoModel: newBlogUserInfoModel(conn, c, opts...),
	}
}

var (
	cachePublicBlogUserInfoUserUuidAndBlogUuidPrefix = "cache:public:blogUserInfo:userUuidAndblogUuid:"
)

func (c *customBlogUserInfoModel) FindOneByUserUuidAndBlogUuid(ctx context.Context, userUuid string, blogUuid string) (*BlogUserInfo, error) {
	publicBlogUserInfoUserUuidAndBlogUuidKey := fmt.Sprintf("%s%v%v", cachePublicBlogUserInfoUserUuidAndBlogUuidPrefix, userUuid, blogUuid)
	var resp BlogUserInfo
	err := c.QueryRowIndexCtx(ctx, &resp, publicBlogUserInfoUserUuidAndBlogUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where user_uuid = $1 and blog_uuid = $2 limit 1", blogUserInfoRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userUuid, blogUuid); err != nil {
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
