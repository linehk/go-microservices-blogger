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
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blogservice"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/config"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/server"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	appUserRepo := model.NewMockAppUserModel(ctrl)
	localeRepo := model.NewMockLocaleModel(ctrl)
	blogService := blogservice.NewMockBlogService(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		AppUserModel: appUserRepo,
		LocaleModel:  localeRepo,
		BlogService:  blogService,
	})
	defer ctrl.Finish()

	userId := uuid.New().String()
	created := time.Now()
	url := "Url"
	userSelfLink := "UserSelfLink"
	displayName := "DisplayName"
	about := "About"
	appUserModel := &model.AppUser{
		Id:          1,
		Uuid:        userId,
		Created:     sql.NullTime{Time: created, Valid: true},
		Url:         sql.NullString{String: url, Valid: true},
		SelfLink:    sql.NullString{String: userSelfLink, Valid: true},
		DisplayName: sql.NullString{String: displayName, Valid: true},
		About:       sql.NullString{String: about, Valid: true},
	}

	localeId := uuid.New().String()
	language := "Language"
	country := "Country"
	variant := "Variant"
	localeModel := &model.Locale{
		Id:          1,
		Uuid:        localeId,
		AppUserUuid: userId,
		Language:    sql.NullString{String: language, Valid: true},
		Country:     sql.NullString{String: country, Valid: true},
		Variant:     sql.NullString{String: variant, Valid: true},
	}

	// UserNotExist
	expectedErr := errcode.Wrap(errcode.UserNotExist)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Get(&user.GetReq{UserId: userId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(&user.GetReq{UserId: userId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// LocaleNotExist
	expectedErr = errcode.Wrap(errcode.LocaleNotExist)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(appUserModel, nil)
	localeRepo.EXPECT().FindOneByAppUserUuid(ctx, userId).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(&user.GetReq{UserId: userId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(appUserModel, nil)
	localeRepo.EXPECT().FindOneByAppUserUuid(ctx, userId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(&user.GetReq{UserId: userId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Service
	listByUserReq := &blog.ListByUserReq{
		UserId: userId,
	}
	expectedErr = errcode.Wrap(errcode.Service)
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(appUserModel, nil)
	localeRepo.EXPECT().FindOneByAppUserUuid(ctx, userId).Return(localeModel, nil)
	blogService.EXPECT().ListByUser(ctx, listByUserReq).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(&user.GetReq{UserId: userId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	blogSelfLink1 := "blogSelfLink1"
	blogSelfLink2 := "blogSelfLink2"
	listByUserResp := &blog.ListByUserResp{
		Items: []*blog.Blog{
			{
				SelfLink: blogSelfLink1,
			},
			{
				SelfLink: blogSelfLink2,
			}},
	}
	expected := &user.User{
		Kind:        "blogger#user",
		Id:          userId,
		Created:     timestamppb.New(created),
		Url:         url,
		SelfLink:    userSelfLink,
		Blogs:       []*user.Blogs{{SelfLink: blogSelfLink1}, {SelfLink: blogSelfLink2}},
		DisplayName: displayName,
		About:       about,
		Locale: &user.Locale{
			Language: language,
			Country:  country,
			Variant:  variant,
		},
	}
	appUserRepo.EXPECT().FindOneByUuid(ctx, userId).Return(appUserModel, nil)
	localeRepo.EXPECT().FindOneByAppUserUuid(ctx, userId).Return(localeModel, nil)
	blogService.EXPECT().ListByUser(ctx, listByUserReq).Return(listByUserResp, nil)
	actual, actualErr = logicService.Get(&user.GetReq{UserId: userId})
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
