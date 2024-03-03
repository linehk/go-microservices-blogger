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

func TestListByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	blogRepo := model.NewMockBlogModel(ctrl)
	postService := postservice.NewMockPostService(ctrl)
	pageService := pageservice.NewMockPageService(ctrl)
	logicService := logic.NewListByUserLogic(ctx, &svc.ServiceContext{
		BlogModel:   blogRepo,
		PostService: postService,
		PageService: pageService,
	})
	defer ctrl.Finish()

	userId := uuid.NewString()
	blogId1 := uuid.NewString()
	blogId2 := uuid.NewString()
	listByUserReq := &blog.ListByUserReq{UserId: userId}

	name := "Name"
	description := "Description"
	published := time.Now()
	updated := time.Now()
	url := "Url"
	selfLink := "SelfLink"
	customMetaData := "CustomMetaData"
	blogModelList := []*model.Blog{
		{
			Id:             1,
			Uuid:           blogId1,
			AppUserUuid:    userId,
			Name:           sql.NullString{String: name, Valid: true},
			Description:    sql.NullString{String: description, Valid: true},
			Published:      sql.NullTime{Time: published, Valid: true},
			Updated:        sql.NullTime{Time: updated, Valid: true},
			Url:            sql.NullString{String: url, Valid: true},
			SelfLink:       sql.NullString{String: selfLink, Valid: true},
			CustomMetaData: sql.NullString{String: customMetaData, Valid: true},
		},
		{
			Id:             2,
			Uuid:           blogId2,
			AppUserUuid:    userId,
			Name:           sql.NullString{String: name, Valid: true},
			Description:    sql.NullString{String: description, Valid: true},
			Published:      sql.NullTime{Time: published, Valid: true},
			Updated:        sql.NullTime{Time: updated, Valid: true},
			Url:            sql.NullString{String: url, Valid: true},
			SelfLink:       sql.NullString{String: selfLink, Valid: true},
			CustomMetaData: sql.NullString{String: customMetaData, Valid: true},
		},
	}

	listPostReq1 := &post.ListReq{
		BlogId: blogId1,
	}
	listPostReq2 := &post.ListReq{
		BlogId: blogId2,
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
	listPageReq1 := &page.ListReq{
		BlogId: blogId1,
	}
	listPageReq2 := &page.ListReq{
		BlogId: blogId2,
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
	expected := &blog.ListByUserResp{
		Kind: "blogger#blogList",
		Items: []*blog.Blog{
			{
				Kind:           "blogger#blog",
				Id:             blogId1,
				Name:           name,
				Description:    description,
				Published:      timestamppb.New(published),
				Updated:        timestamppb.New(updated),
				Url:            url,
				SelfLink:       selfLink,
				Posts:          []*blog.Posts{{TotalItems: postTotalItems, SelfLink: postSelfLink1}, {TotalItems: postTotalItems, SelfLink: postSelfLink2}},
				Pages:          []*blog.Pages{{TotalItems: pageTotalItems, SelfLink: pageSelfLink1}, {TotalItems: pageTotalItems, SelfLink: pageSelfLink2}},
				CustomMetaData: customMetaData,
			}, {
				Kind:           "blogger#blog",
				Id:             blogId2,
				Name:           name,
				Description:    description,
				Published:      timestamppb.New(published),
				Updated:        timestamppb.New(updated),
				Url:            url,
				SelfLink:       selfLink,
				Posts:          []*blog.Posts{{TotalItems: postTotalItems, SelfLink: postSelfLink1}, {TotalItems: postTotalItems, SelfLink: postSelfLink2}},
				Pages:          []*blog.Pages{{TotalItems: pageTotalItems, SelfLink: pageSelfLink1}, {TotalItems: pageTotalItems, SelfLink: pageSelfLink2}},
				CustomMetaData: customMetaData,
			},
		},
	}

	// BlogNotExist
	expectedErr := errcode.Wrap(errcode.BlogNotExist)
	blogRepo.EXPECT().ListByAppUserUuid(ctx, userId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.ListByUser(listByUserReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	blogRepo.EXPECT().ListByAppUserUuid(ctx, userId).Return(nil, expectedErr)
	actual, actualErr = logicService.ListByUser(listByUserReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Post Service
	expectedErr = errcode.Wrap(errcode.Service)
	blogRepo.EXPECT().ListByAppUserUuid(ctx, userId).Return(blogModelList, nil)
	postService.EXPECT().List(ctx, listPostReq1).Return(nil, expectedErr)

	actual, actualErr = logicService.ListByUser(listByUserReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Page Service
	expectedErr = errcode.Wrap(errcode.Service)
	blogRepo.EXPECT().ListByAppUserUuid(ctx, userId).Return(blogModelList, nil)
	postService.EXPECT().List(ctx, listPostReq1).Return(listPostResp, nil)

	pageService.EXPECT().List(ctx, listPageReq1).Return(nil, expectedErr)
	actual, actualErr = logicService.ListByUser(listByUserReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	blogRepo.EXPECT().ListByAppUserUuid(ctx, userId).Return(blogModelList, nil)
	postService.EXPECT().List(ctx, listPostReq1).Return(listPostResp, nil)
	postService.EXPECT().List(ctx, listPostReq2).Return(listPostResp, nil)
	pageService.EXPECT().List(ctx, listPageReq1).Return(listPageResp, nil)
	pageService.EXPECT().List(ctx, listPageReq2).Return(listPageResp, nil)
	actual, actualErr = logicService.ListByUser(listByUserReq)
	assert.Equal(t, actual, expected)
	assert.Nil(t, actualErr)
}
