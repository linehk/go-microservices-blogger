package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
	postmodel "github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageRepo := model.NewMockPageModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		PageModel:   pageRepo,
		AuthorModel: authorRepo,
		ImageModel:  imageRepo,
	})
	defer ctrl.Finish()

	blogUuid := uuid.NewString()
	pageUuid := uuid.NewString()
	getReq := &page.GetReq{
		BlogId: blogUuid,
		PageId: pageUuid,
	}

	status := "Status"
	published := time.Now()
	updated := time.Now()
	pageUrl := "Url"
	selfLink := "SelfLink"
	title := "Title"
	content := "Content"
	pageModel := &model.Page{
		Id:        1,
		Uuid:      pageUuid,
		BlogUuid:  sql.NullString{String: blogUuid, Valid: true},
		Status:    sql.NullString{String: status, Valid: true},
		Published: sql.NullTime{Time: published, Valid: true},
		Updated:   sql.NullTime{Time: updated, Valid: true},
		Url:       sql.NullString{String: pageUrl, Valid: true},
		SelfLink:  sql.NullString{String: selfLink, Valid: true},
		Title:     sql.NullString{String: title, Valid: true},
		Content:   sql.NullString{String: content, Valid: true},
	}

	authorUuid := uuid.NewString()
	displayName := "DisplayName"
	authorUrl := "Url"
	authorModel := &postmodel.Author{
		Id:          1,
		Uuid:        authorUuid,
		PostUuid:    "",
		PageUuid:    pageUuid,
		CommentUuid: "",
		DisplayName: sql.NullString{String: displayName, Valid: true},
		Url:         sql.NullString{String: authorUrl, Valid: true},
	}

	imageUuid := uuid.NewString()
	imageUrl := "imageUrl"
	imageModel := &postmodel.Image{
		Id:         1,
		Uuid:       imageUuid,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid,
		Url:        sql.NullString{String: imageUrl, Valid: true},
	}

	expected := &page.Page{
		Kind:      "blogger#page",
		Id:        pageUuid,
		Status:    status,
		Blog:      &page.Blog{Id: blogUuid},
		Published: timestamppb.New(published),
		Updated:   timestamppb.New(updated),
		Url:       pageUrl,
		SelfLink:  selfLink,
		Title:     title,
		Content:   content,
		Author: &page.Author{
			Id:          authorUuid,
			DisplayName: displayName,
			Url:         authorUrl,
			Image:       &page.Image{Url: imageUrl},
		},
	}

	// PageNotExist
	expectedErr := errcode.Wrap(errcode.PageNotExist)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorNotExist
	expectedErr = errcode.Wrap(errcode.AuthorNotExist)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// FindOneByPageUuid Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// ImageNotExist
	expectedErr = errcode.Wrap(errcode.ImageNotExist)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// FindOneByAuthorUuid Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(imageModel, nil)
	actual, actualErr = logicService.Get(getReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
