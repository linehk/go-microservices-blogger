package main

import (
	"flag"
	"fmt"

	"github.com/linehk/go-blogger/service/user_center/rpc/internal/config"
	"github.com/linehk/go-blogger/service/user_center/rpc/internal/server"
	"github.com/linehk/go-blogger/service/user_center/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user_center/rpc/user_center"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user_center.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_center.RegisterUserCenterServer(grpcServer, server.NewUserCenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
