package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/pageservice"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/postservice"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	blogRepo := model.NewMockBlogModel(ctrl)
	postService := postservice.NewMockPostService(ctrl)
	pageService := pageservice.NewMockPageService(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		BlogModel:   blogRepo,
		PostService: postService,
		PageService: pageService,
	})
	defer ctrl.Finish()

	blogId := uuid.New().String()

	getReq := &blog.GetReq{BlogId: blogId}

	userId := uuid.New().String()
	name := "Name"
	description := "Description"
	published := time.Now()
	updated := time.Now()
	url := "Url"
	selfLink := "SelfLink"
	customMetaData := "CustomMetaData"
	blogModel := &model.Blog{
		Id:             1,
		Uuid:           blogId,
		AppUserUuid:    userId,
		Name:           sql.NullString{String: name, Valid: true},
		Description:    sql.NullString{String: description, Valid: true},
		Published:      sql.NullTime{Time: published, Valid: true},
		Updated:        sql.NullTime{Time: updated, Valid: true},
		Url:            sql.NullString{String: url, Valid: true},
		SelfLink:       sql.NullString{String: selfLink, Valid: true},
		CustomMetaData: sql.NullString{String: customMetaData, Valid: true},
	}

	listPostReq := &post.ListReq{
		BlogId: blogId,
	}
	postSelfLink1 := "postSelfLink1"
	postSelfLink2 := "postSelfLink2"
	listPostResp := &post.ListResp{
		Kind: "blogger#post",
		Items: []*post.Post{{
			SelfLink: postSelfLink1,
		}, {
			SelfLink: postSelfLink2,
		}},
	}

	pageSelfLink1 := "pageSelfLink1"
	pageSelfLink2 := "pageSelfLink2"
	listPageReq := &page.ListReq{
		BlogId: blogId,
	}
	listPageResp := &page.ListResp{
		Kind: "blogger#page",
		Items: []*page.Page{{
			SelfLink: pageSelfLink1,
		}, {
			SelfLink: pageSelfLink2,
		}},
	}

	postTotalItems := "2"
	pageTotalItems := "2"
	expected := &blog.Blog{
		Kind:           "blogger#blog",
		Id:             blogId,
		Name:           name,
		Description:    description,
		Published:      timestamppb.New(published),
		Updated:        timestamppb.New(updated),
		Url:            url,
		SelfLink:       selfLink,
		Posts:          []*blog.Posts{{TotalItems: postTotalItems, SelfLink: postSelfLink1}, {TotalItems: postTotalItems, SelfLink: postSelfLink2}},
		Pages:          []*blog.Pages{{TotalItems: pageTotalItems, SelfLink: pageSelfLink1}, {TotalItems: pageTotalItems, SelfLink: pageSelfLink2}},
		CustomMetaData: customMetaData,
	}

	// BlogNotExist
	expectedErr := errcode.Wrap(errcode.BlogNotExist)
	blogRepo.EXPECT().FindOneByUuid(ctx, blogId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	blogRepo.EXPECT().FindOneByUuid(ctx, blogId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Post Service
	expectedErr = errcode.Wrap(errcode.Service)
	blogRepo.EXPECT().FindOneByUuid(ctx, blogId).Return(blogModel, nil)
	postService.EXPECT().List(ctx, listPostReq).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Page Service
	expectedErr = errcode.Wrap(errcode.Service)
	blogRepo.EXPECT().FindOneByUuid(ctx, blogId).Return(blogModel, nil)
	postService.EXPECT().List(ctx, listPostReq).Return(listPostResp, nil)
	pageService.EXPECT().List(ctx, listPageReq).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	blogRepo.EXPECT().FindOneByUuid(ctx, blogId).Return(blogModel, nil)
	postService.EXPECT().List(ctx, listPostReq).Return(listPostResp, nil)
	pageService.EXPECT().List(ctx, listPageReq).Return(listPageResp, nil)
	actual, actualErr = logicService.Get(getReq)
	assert.Equal(t, actual, expected)
	assert.Nil(t, actualErr)
}
