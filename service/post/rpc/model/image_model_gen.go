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
	imageFieldNames          = builder.RawFieldNames(&Image{}, true)
	imageRows                = strings.Join(imageFieldNames, ",")
	imageRowsExpectAutoSet   = strings.Join(stringx.Remove(imageFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	imageRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(imageFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicImageIdPrefix         = "cache:public:image:id:"
	cachePublicImageAuthorUuidPrefix = "cache:public:image:authorUuid:"
	cachePublicImagePostUuidPrefix   = "cache:public:image:postUuid:"
	cachePublicImageUuidPrefix       = "cache:public:image:uuid:"
)

type (
	imageModel interface {
		Insert(ctx context.Context, data *Image) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Image, error)
		FindOneByAuthorUuid(ctx context.Context, authorUuid string) (*Image, error)
		FindOneByPostUuid(ctx context.Context, postUuid string) (*Image, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Image, error)
		Update(ctx context.Context, data *Image) error
		Delete(ctx context.Context, id int64) error
	}

	defaultImageModel struct {
		sqlc.CachedConn
		table string
	}

	Image struct {
		Id         int64          `db:"id"`
		Uuid       string         `db:"uuid"`
		PostUuid   string         `db:"post_uuid"`
		AuthorUuid string         `db:"author_uuid"`
		Url        sql.NullString `db:"url"`
	}
)

func newImageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultImageModel {
	return &defaultImageModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."image"`,
	}
}

func (m *defaultImageModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicImageAuthorUuidKey := fmt.Sprintf("%s%v", cachePublicImageAuthorUuidPrefix, data.AuthorUuid)
	publicImageIdKey := fmt.Sprintf("%s%v", cachePublicImageIdPrefix, id)
	publicImagePostUuidKey := fmt.Sprintf("%s%v", cachePublicImagePostUuidPrefix, data.PostUuid)
	publicImageUuidKey := fmt.Sprintf("%s%v", cachePublicImageUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicImageAuthorUuidKey, publicImageIdKey, publicImagePostUuidKey, publicImageUuidKey)
	return err
}

func (m *defaultImageModel) FindOne(ctx context.Context, id int64) (*Image, error) {
	publicImageIdKey := fmt.Sprintf("%s%v", cachePublicImageIdPrefix, id)
	var resp Image
	err := m.QueryRowCtx(ctx, &resp, publicImageIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", imageRows, m.table)
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

func (m *defaultImageModel) FindOneByAuthorUuid(ctx context.Context, authorUuid string) (*Image, error) {
	publicImageAuthorUuidKey := fmt.Sprintf("%s%v", cachePublicImageAuthorUuidPrefix, authorUuid)
	var resp Image
	err := m.QueryRowIndexCtx(ctx, &resp, publicImageAuthorUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where author_uuid = $1 limit 1", imageRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, authorUuid); err != nil {
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

func (m *defaultImageModel) FindOneByPostUuid(ctx context.Context, postUuid string) (*Image, error) {
	publicImagePostUuidKey := fmt.Sprintf("%s%v", cachePublicImagePostUuidPrefix, postUuid)
	var resp Image
	err := m.QueryRowIndexCtx(ctx, &resp, publicImagePostUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where post_uuid = $1 limit 1", imageRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, postUuid); err != nil {
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

func (m *defaultImageModel) FindOneByUuid(ctx context.Context, uuid string) (*Image, error) {
	publicImageUuidKey := fmt.Sprintf("%s%v", cachePublicImageUuidPrefix, uuid)
	var resp Image
	err := m.QueryRowIndexCtx(ctx, &resp, publicImageUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", imageRows, m.table)
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

func (m *defaultImageModel) Insert(ctx context.Context, data *Image) (sql.Result, error) {
	publicImageAuthorUuidKey := fmt.Sprintf("%s%v", cachePublicImageAuthorUuidPrefix, data.AuthorUuid)
	publicImageIdKey := fmt.Sprintf("%s%v", cachePublicImageIdPrefix, data.Id)
	publicImagePostUuidKey := fmt.Sprintf("%s%v", cachePublicImagePostUuidPrefix, data.PostUuid)
	publicImageUuidKey := fmt.Sprintf("%s%v", cachePublicImageUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, imageRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.PostUuid, data.AuthorUuid, data.Url)
	}, publicImageAuthorUuidKey, publicImageIdKey, publicImagePostUuidKey, publicImageUuidKey)
	return ret, err
}

func (m *defaultImageModel) Update(ctx context.Context, newData *Image) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicImageAuthorUuidKey := fmt.Sprintf("%s%v", cachePublicImageAuthorUuidPrefix, data.AuthorUuid)
	publicImageIdKey := fmt.Sprintf("%s%v", cachePublicImageIdPrefix, data.Id)
	publicImagePostUuidKey := fmt.Sprintf("%s%v", cachePublicImagePostUuidPrefix, data.PostUuid)
	publicImageUuidKey := fmt.Sprintf("%s%v", cachePublicImageUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, imageRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.PostUuid, newData.AuthorUuid, newData.Url)
	}, publicImageAuthorUuidKey, publicImageIdKey, publicImagePostUuidKey, publicImageUuidKey)
	return err
}

func (m *defaultImageModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicImageIdPrefix, primary)
}

func (m *defaultImageModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", imageRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultImageModel) tableName() string {
	return m.table
}