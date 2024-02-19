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
	labelFieldNames          = builder.RawFieldNames(&Label{}, true)
	labelRows                = strings.Join(labelFieldNames, ",")
	labelRowsExpectAutoSet   = strings.Join(stringx.Remove(labelFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	labelRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(labelFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicLabelIdPrefix       = "cache:public:label:id:"
	cachePublicLabelPostUuidPrefix = "cache:public:label:postUuid:"
	cachePublicLabelUuidPrefix     = "cache:public:label:uuid:"
)

type (
	labelModel interface {
		Insert(ctx context.Context, data *Label) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Label, error)
		FindOneByPostUuid(ctx context.Context, postUuid string) (*Label, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Label, error)
		Update(ctx context.Context, data *Label) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLabelModel struct {
		sqlc.CachedConn
		table string
	}

	Label struct {
		Id         int64          `db:"id"`
		Uuid       string         `db:"uuid"`
		PostUuid   string         `db:"post_uuid"`
		LabelValue sql.NullString `db:"label_value"`
	}
)

func newLabelModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLabelModel {
	return &defaultLabelModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."label"`,
	}
}

func (m *defaultLabelModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicLabelIdKey := fmt.Sprintf("%s%v", cachePublicLabelIdPrefix, id)
	publicLabelPostUuidKey := fmt.Sprintf("%s%v", cachePublicLabelPostUuidPrefix, data.PostUuid)
	publicLabelUuidKey := fmt.Sprintf("%s%v", cachePublicLabelUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicLabelIdKey, publicLabelPostUuidKey, publicLabelUuidKey)
	return err
}

func (m *defaultLabelModel) FindOne(ctx context.Context, id int64) (*Label, error) {
	publicLabelIdKey := fmt.Sprintf("%s%v", cachePublicLabelIdPrefix, id)
	var resp Label
	err := m.QueryRowCtx(ctx, &resp, publicLabelIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", labelRows, m.table)
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

func (m *defaultLabelModel) FindOneByPostUuid(ctx context.Context, postUuid string) (*Label, error) {
	publicLabelPostUuidKey := fmt.Sprintf("%s%v", cachePublicLabelPostUuidPrefix, postUuid)
	var resp Label
	err := m.QueryRowIndexCtx(ctx, &resp, publicLabelPostUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where post_uuid = $1 limit 1", labelRows, m.table)
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

func (m *defaultLabelModel) FindOneByUuid(ctx context.Context, uuid string) (*Label, error) {
	publicLabelUuidKey := fmt.Sprintf("%s%v", cachePublicLabelUuidPrefix, uuid)
	var resp Label
	err := m.QueryRowIndexCtx(ctx, &resp, publicLabelUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", labelRows, m.table)
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

func (m *defaultLabelModel) Insert(ctx context.Context, data *Label) (sql.Result, error) {
	publicLabelIdKey := fmt.Sprintf("%s%v", cachePublicLabelIdPrefix, data.Id)
	publicLabelPostUuidKey := fmt.Sprintf("%s%v", cachePublicLabelPostUuidPrefix, data.PostUuid)
	publicLabelUuidKey := fmt.Sprintf("%s%v", cachePublicLabelUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, labelRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.PostUuid, data.LabelValue)
	}, publicLabelIdKey, publicLabelPostUuidKey, publicLabelUuidKey)
	return ret, err
}

func (m *defaultLabelModel) Update(ctx context.Context, newData *Label) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicLabelIdKey := fmt.Sprintf("%s%v", cachePublicLabelIdPrefix, data.Id)
	publicLabelPostUuidKey := fmt.Sprintf("%s%v", cachePublicLabelPostUuidPrefix, data.PostUuid)
	publicLabelUuidKey := fmt.Sprintf("%s%v", cachePublicLabelUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, labelRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.PostUuid, newData.LabelValue)
	}, publicLabelIdKey, publicLabelPostUuidKey, publicLabelUuidKey)
	return err
}

func (m *defaultLabelModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicLabelIdPrefix, primary)
}

func (m *defaultLabelModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", labelRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLabelModel) tableName() string {
	return m.table
}