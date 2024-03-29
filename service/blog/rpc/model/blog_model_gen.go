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
	blogFieldNames          = builder.RawFieldNames(&Blog{}, true)
	blogRows                = strings.Join(blogFieldNames, ",")
	blogRowsExpectAutoSet   = strings.Join(stringx.Remove(blogFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	blogRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(blogFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicBlogIdPrefix   = "cache:public:blog:id:"
	cachePublicBlogUuidPrefix = "cache:public:blog:uuid:"
)

type (
	blogModel interface {
		Insert(ctx context.Context, data *Blog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Blog, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Blog, error)
		Update(ctx context.Context, data *Blog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBlogModel struct {
		sqlc.CachedConn
		table string
	}

	Blog struct {
		Id             int64          `db:"id"`
		Uuid           string         `db:"uuid"`
		AppUserUuid    sql.NullString `db:"app_user_uuid"`
		Name           sql.NullString `db:"name"`
		Description    sql.NullString `db:"description"`
		Published      sql.NullTime   `db:"published"`
		Updated        sql.NullTime   `db:"updated"`
		Url            sql.NullString `db:"url"`
		SelfLink       sql.NullString `db:"self_link"`
		CustomMetaData sql.NullString `db:"custom_meta_data"`
	}
)

func newBlogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultBlogModel {
	return &defaultBlogModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."blog"`,
	}
}

func (m *defaultBlogModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicBlogIdKey := fmt.Sprintf("%s%v", cachePublicBlogIdPrefix, id)
	publicBlogUuidKey := fmt.Sprintf("%s%v", cachePublicBlogUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicBlogIdKey, publicBlogUuidKey)
	return err
}

func (m *defaultBlogModel) FindOne(ctx context.Context, id int64) (*Blog, error) {
	publicBlogIdKey := fmt.Sprintf("%s%v", cachePublicBlogIdPrefix, id)
	var resp Blog
	err := m.QueryRowCtx(ctx, &resp, publicBlogIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", blogRows, m.table)
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

func (m *defaultBlogModel) FindOneByUuid(ctx context.Context, uuid string) (*Blog, error) {
	publicBlogUuidKey := fmt.Sprintf("%s%v", cachePublicBlogUuidPrefix, uuid)
	var resp Blog
	err := m.QueryRowIndexCtx(ctx, &resp, publicBlogUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", blogRows, m.table)
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

func (m *defaultBlogModel) Insert(ctx context.Context, data *Blog) (sql.Result, error) {
	publicBlogIdKey := fmt.Sprintf("%s%v", cachePublicBlogIdPrefix, data.Id)
	publicBlogUuidKey := fmt.Sprintf("%s%v", cachePublicBlogUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", m.table, blogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.AppUserUuid, data.Name, data.Description, data.Published, data.Updated, data.Url, data.SelfLink, data.CustomMetaData)
	}, publicBlogIdKey, publicBlogUuidKey)
	return ret, err
}

func (m *defaultBlogModel) Update(ctx context.Context, newData *Blog) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicBlogIdKey := fmt.Sprintf("%s%v", cachePublicBlogIdPrefix, data.Id)
	publicBlogUuidKey := fmt.Sprintf("%s%v", cachePublicBlogUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, blogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.AppUserUuid, newData.Name, newData.Description, newData.Published, newData.Updated, newData.Url, newData.SelfLink, newData.CustomMetaData)
	}, publicBlogIdKey, publicBlogUuidKey)
	return err
}

func (m *defaultBlogModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicBlogIdPrefix, primary)
}

func (m *defaultBlogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", blogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBlogModel) tableName() string {
	return m.table
}
