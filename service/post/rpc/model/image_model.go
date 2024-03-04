package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ImageModel = (*customImageModel)(nil)

//go:generate mockgen -destination=./mock_image_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model ImageModel
type (
	// ImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImageModel.
	ImageModel interface {
		imageModel
		ListByPostUuid(ctx context.Context, postUuid string) ([]*Image, error)
	}

	customImageModel struct {
		*defaultImageModel
	}
)

// NewImageModel returns a model for the database table.
func NewImageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ImageModel {
	return &customImageModel{
		defaultImageModel: newImageModel(conn, c, opts...),
	}
}

var (
	cachePublicImageListPostUuidPrefix = "cache:public:image:list:postUuid:"
)

func (c *customImageModel) ListByPostUuid(ctx context.Context, postUuid string) ([]*Image, error) {
	publicImageListPostUuidKey := fmt.Sprintf("%s%v", cachePublicImageListPostUuidPrefix, postUuid)
	var resp []*Image
	err := c.QueryRowIndexCtx(ctx, &resp, publicImageListPostUuidKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where post_uuid = $1", imageRows, c.table)
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
