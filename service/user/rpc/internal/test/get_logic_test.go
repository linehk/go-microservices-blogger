package test

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"log"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/server"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"
	"github.com/stretchr/testify/require"
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

func newCtrl(t *testing.T) (
	*gomock.Controller,
	context.Context,
	*logic.GetLogic,
	*model.MockAppUserModel,
	*model.MockLocaleModel) {

	ctrl := gomock.NewController(t)
	ctx := context.Background()
	appUserRepo := model.NewMockAppUserModel(ctrl)
	localeRepo := model.NewMockLocaleModel(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		AppUserModel: appUserRepo,
		LocaleModel:  localeRepo,
	})
	return ctrl, ctx, logicService, appUserRepo, localeRepo
}

func TestGetUserNotExist(t *testing.T) {
	ctrl, ctx, logicService, appUserRepo, _ := newCtrl(t)
	defer ctrl.Finish()

	userId := uuid.New().String()

	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(nil, model.ErrNotFound)

	_, actual := logicService.Get(&user.GetReq{UserId: userId})
	expected := errcode.Wrap(errcode.UserNotExist)
	require.Equal(t, expected, actual)
}

func TestGetAppUserDatabase(t *testing.T) {
	ctrl, ctx, logicService, appUserRepo, _ := newCtrl(t)
	defer ctrl.Finish()

	userId := uuid.New().String()

	expected := errcode.Wrap(errcode.Database)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(nil, expected)

	_, actual := logicService.Get(&user.GetReq{UserId: userId})
	require.Equal(t, actual, expected)
}

func TestGetLocaleNotExist(t *testing.T) {
	ctrl, ctx, logicService, appUserRepo, localeRepo := newCtrl(t)
	defer ctrl.Finish()

	userId := uuid.New().String()

	appUserModel := &model.AppUser{
		Id:          1,
		Uuid:        userId,
		Created:     sql.NullTime{Time: time.Now(), Valid: true},
		Url:         sql.NullString{String: "Url", Valid: true},
		SelfLink:    sql.NullString{String: "SelfLink", Valid: true},
		DisplayName: sql.NullString{String: "DisplayName", Valid: true},
		About:       sql.NullString{String: "About", Valid: true},
	}
	
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(appUserModel, nil)
	localeRepo.EXPECT().FindOneByAppUserUuid(ctx, userId).Return(nil, model.ErrNotFound)

	expected := errcode.Wrap(errcode.LocaleNotExist)

	_, actual := logicService.Get(&user.GetReq{UserId: userId})
	require.Equal(t, expected, actual)
}
