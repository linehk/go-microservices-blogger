package test

import (
	"context"
	"errors"
	"flag"
	"net"
	"testing"

	"github.com/emicklei/go-restful/v3/log"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/server"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var configFile = flag.String("f", "../../etc/user_test.yaml", "the config file")

func serverStart(ctx context.Context) (user.UserServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	serviceContext := svc.NewServiceContext(c)
	user.RegisterUserServiceServer(baseServer, server.NewUserServiceServer(serviceContext))

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := user.NewUserServiceClient(conn)

	return client, closer
}

func TGet(t *testing.T) {
	ctx := context.Background()

	client, closer := serverStart(ctx)
	defer closer()

	tests := []struct {
		input   *user.GetReq
		want    *user.User
		wantErr error
	}{
		{
			&user.GetReq{UserId: "07d62983-3b68-44a6-ae95-0f297be74606"},
			&user.User{
				Kind:        "blogger#user",
				Id:          "07d62983-3b68-44a6-ae95-0f297be74606",
				Created:     nil,
				Url:         "",
				SelfLink:    "",
				Blogs:       nil,
				DisplayName: "",
				About:       "",
				Locale:      nil,
			},
			nil,
		},
	}

	for _, tt := range tests {
		got, gotErr := client.Get(ctx, tt.input)
		if !errors.Is(gotErr, tt.wantErr) {
			t.Errorf("got %v, want %v", gotErr, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
