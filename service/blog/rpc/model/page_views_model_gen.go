// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pageViewsFieldNames          = builder.RawFieldNames(&PageViews{}, true)
	pageViewsRows                = strings.Join(pageViewsFieldNames, ",")
	pageViewsRowsExpectAutoSet   = strings.Join(stringx.Remove(pageViewsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	pageViewsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(pageViewsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicPageViewsIdPrefix       = "cache:public:pageViews:id:"
	cachePublicPageViewsBlogUuidPrefix = "cache:public:pageViews:blogUuid:"
	cachePublicPageViewsUuidPrefix     = "cache:public:pageViews:uuid:"
)

type (
	pageViewsModel interface {
		Insert(ctx context.Context, data *PageViews) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PageViews, error)
		FindOneByBlogUuid(ctx context.Context, blogUuid string) (*PageViews, error)
		FindOneByUuid(ctx context.Context, uuid string) (*PageViews, error)
		Update(ctx context.Context, data *PageViews) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPageViewsModel struct {
		sqlc.CachedConn
		table string
	}

	PageViews struct {
		Id       int64         `db:"id"`
		Uuid     string        `db:"uuid"`
		BlogUuid string        `db:"blog_uuid"`
		Count    sql.NullInt64 `db:"count"`
	}
)

func newPageViewsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPageViewsModel {
	return &defaultPageViewsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."page_views"`,
	}
}

func (m *defaultPageViewsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicPageViewsBlogUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsBlogUuidPrefix, data.BlogUuid)
	publicPageViewsIdKey := fmt.Sprintf("%s%v", cachePublicPageViewsIdPrefix, id)
	publicPageViewsUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicPageViewsBlogUuidKey, publicPageViewsIdKey, publicPageViewsUuidKey)
	return err
}

func (m *defaultPageViewsModel) FindOne(ctx context.Context, id int64) (*PageViews, error) {
	publicPageViewsIdKey := fmt.Sprintf("%s%v", cachePublicPageViewsIdPrefix, id)
	var resp PageViews
	err := m.QueryRowCtx(ctx, &resp, publicPageViewsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", pageViewsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPageViewsModel) FindOneByBlogUuid(ctx context.Context, blogUuid string) (*PageViews, error) {
	publicPageViewsBlogUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsBlogUuidPrefix, blogUuid)
	var resp PageViews
	err := m.QueryRowIndexCtx(ctx, &resp, publicPageViewsBlogUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where blog_uuid = $1 limit 1", pageViewsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, blogUuid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPageViewsModel) FindOneByUuid(ctx context.Context, uuid string) (*PageViews, error) {
	publicPageViewsUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsUuidPrefix, uuid)
	var resp PageViews
	err := m.QueryRowIndexCtx(ctx, &resp, publicPageViewsUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", pageViewsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uuid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPageViewsModel) Insert(ctx context.Context, data *PageViews) (sql.Result, error) {
	publicPageViewsBlogUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsBlogUuidPrefix, data.BlogUuid)
	publicPageViewsIdKey := fmt.Sprintf("%s%v", cachePublicPageViewsIdPrefix, data.Id)
	publicPageViewsUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, pageViewsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.BlogUuid, data.Count)
	}, publicPageViewsBlogUuidKey, publicPageViewsIdKey, publicPageViewsUuidKey)
	return ret, err
}

func (m *defaultPageViewsModel) Update(ctx context.Context, newData *PageViews) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicPageViewsBlogUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsBlogUuidPrefix, data.BlogUuid)
	publicPageViewsIdKey := fmt.Sprintf("%s%v", cachePublicPageViewsIdPrefix, data.Id)
	publicPageViewsUuidKey := fmt.Sprintf("%s%v", cachePublicPageViewsUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, pageViewsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.BlogUuid, newData.Count)
	}, publicPageViewsBlogUuidKey, publicPageViewsIdKey, publicPageViewsUuidKey)
	return err
}

func (m *defaultPageViewsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicPageViewsIdPrefix, primary)
}

func (m *defaultPageViewsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", pageViewsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPageViewsModel) tableName() string {
	return m.table
}
