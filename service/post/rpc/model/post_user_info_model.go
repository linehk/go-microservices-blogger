package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostUserInfoModel = (*customPostUserInfoModel)(nil)

//go:generate mockgen -destination=./mock_post_user_info_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model PostUserInfoModel
type (
	// PostUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostUserInfoModel.
	PostUserInfoModel interface {
		postUserInfoModel
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
