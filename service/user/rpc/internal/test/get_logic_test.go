package test

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"testing"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/server"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/conf"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var configFile = flag.String("f", "../../etc/user_test.yaml", "the config file")

func startServer(ctx context.Context) (user.UserServiceClient, func()) {
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

func IntegrationTestGet(t *testing.T) {
	ctx := context.Background()

	client, closer := startServer(ctx)
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

func newCtrl(t *testing.T) (*gomock.Controller, *model.MockAppUserModel, context.Context, *logic.GetLogic) {
	ctrl := gomock.NewController(t)
	repo := model.NewMockAppUserModel(ctrl)
	ctx := context.Background()
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		AppUserModel: repo,
	})
	return ctrl, repo, ctx, logicService
}

func TestGetUserNotExist(t *testing.T) {
	ctrl, repo, ctx, logicService := newCtrl(t)
	defer ctrl.Finish()

	userId := uuid.New().String()

	repo.EXPECT().FindOneByUuid(ctx, userId).Return(nil, model.ErrNotFound)

	_, gotErr := logicService.Get(&user.GetReq{UserId: userId})
	wantErr := errcode.Wrap(errcode.UserNotExist)
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("got %v, want %v", gotErr, wantErr)
	}
}
