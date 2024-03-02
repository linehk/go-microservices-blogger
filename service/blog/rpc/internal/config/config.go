package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache    cache.CacheConf
	PostConf zrpc.RpcClientConf
	PageConf zrpc.RpcClientConf
}
