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
	postFieldNames          = builder.RawFieldNames(&Post{}, true)
	postRows                = strings.Join(postFieldNames, ",")
	postRowsExpectAutoSet   = strings.Join(stringx.Remove(postFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	postRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(postFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicPostIdPrefix   = "cache:public:post:id:"
	cachePublicPostUrlPrefix  = "cache:public:post:url:"
	cachePublicPostUuidPrefix = "cache:public:post:uuid:"
)

type (
	postModel interface {
		Insert(ctx context.Context, data *Post) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Post, error)
		FindOneByUrl(ctx context.Context, url string) (*Post, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Post, error)
		Update(ctx context.Context, data *Post) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPostModel struct {
		sqlc.CachedConn
		table string
	}

	Post struct {
		Id             int64          `db:"id"`
		Uuid           string         `db:"uuid"`
		BlogUuid       sql.NullString `db:"blog_uuid"`
		Published      sql.NullTime   `db:"published"`
		Updated        sql.NullTime   `db:"updated"`
		Url            string         `db:"url"`
		SelfLink       sql.NullString `db:"self_link"`
		Title          sql.NullString `db:"title"`
		TitleLink      sql.NullString `db:"title_link"`
		Content        sql.NullString `db:"content"`
		CustomMetaData sql.NullString `db:"custom_meta_data"`
		Status         sql.NullString `db:"status"`
	}
)

func newPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPostModel {
	return &defaultPostModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."post"`,
	}
}

func (m *defaultPostModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicPostIdKey := fmt.Sprintf("%s%v", cachePublicPostIdPrefix, id)
	publicPostUrlKey := fmt.Sprintf("%s%v", cachePublicPostUrlPrefix, data.Url)
	publicPostUuidKey := fmt.Sprintf("%s%v", cachePublicPostUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicPostIdKey, publicPostUrlKey, publicPostUuidKey)
	return err
}

func (m *defaultPostModel) FindOne(ctx context.Context, id int64) (*Post, error) {
	publicPostIdKey := fmt.Sprintf("%s%v", cachePublicPostIdPrefix, id)
	var resp Post
	err := m.QueryRowCtx(ctx, &resp, publicPostIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", postRows, m.table)
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

func (m *defaultPostModel) FindOneByUrl(ctx context.Context, url string) (*Post, error) {
	publicPostUrlKey := fmt.Sprintf("%s%v", cachePublicPostUrlPrefix, url)
	var resp Post
	err := m.QueryRowIndexCtx(ctx, &resp, publicPostUrlKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where url = $1 limit 1", postRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, url); err != nil {
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

func (m *defaultPostModel) FindOneByUuid(ctx context.Context, uuid string) (*Post, error) {
	publicPostUuidKey := fmt.Sprintf("%s%v", cachePublicPostUuidPrefix, uuid)
	var resp Post
	err := m.QueryRowIndexCtx(ctx, &resp, publicPostUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", postRows, m.table)
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

func (m *defaultPostModel) Insert(ctx context.Context, data *Post) (sql.Result, error) {
	publicPostIdKey := fmt.Sprintf("%s%v", cachePublicPostIdPrefix, data.Id)
	publicPostUrlKey := fmt.Sprintf("%s%v", cachePublicPostUrlPrefix, data.Url)
	publicPostUuidKey := fmt.Sprintf("%s%v", cachePublicPostUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", m.table, postRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.BlogUuid, data.Published, data.Updated, data.Url, data.SelfLink, data.Title, data.TitleLink, data.Content, data.CustomMetaData, data.Status)
	}, publicPostIdKey, publicPostUrlKey, publicPostUuidKey)
	return ret, err
}

func (m *defaultPostModel) Update(ctx context.Context, newData *Post) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicPostIdKey := fmt.Sprintf("%s%v", cachePublicPostIdPrefix, data.Id)
	publicPostUrlKey := fmt.Sprintf("%s%v", cachePublicPostUrlPrefix, data.Url)
	publicPostUuidKey := fmt.Sprintf("%s%v", cachePublicPostUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, postRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.BlogUuid, newData.Published, newData.Updated, newData.Url, newData.SelfLink, newData.Title, newData.TitleLink, newData.Content, newData.CustomMetaData, newData.Status)
	}, publicPostIdKey, publicPostUrlKey, publicPostUuidKey)
	return err
}

func (m *defaultPostModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicPostIdPrefix, primary)
}

func (m *defaultPostModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", postRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPostModel) tableName() string {
	return m.table
}
