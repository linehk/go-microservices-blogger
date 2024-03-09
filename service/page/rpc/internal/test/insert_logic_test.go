package test

import (
	"context"
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

func TestInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageRepo := model.NewMockPageModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewInsertLogic(ctx, &svc.ServiceContext{
		PageModel:   pageRepo,
		AuthorModel: authorRepo,
		ImageModel:  imageRepo,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	status := "Status"
	published := time.Now()
	updated := time.Now()
	pageUrl := "Url"
	selfLink := "SelfLink"
	title := "Title"
	content := "Content"

	displayName := "DisplayName"
	authorUrl := "Url"

	imageUrl := "imageUrl"

	insertReq := &page.InsertReq{
		BlogId: blogId,
		Page: &page.Page{
			Status:    status,
			Blog:      &page.Blog{Id: blogId},
			Published: timestamppb.New(published),
			Updated:   timestamppb.New(updated),
			Url:       pageUrl,
			SelfLink:  selfLink,
			Title:     title,
			Content:   content,
			Author: &page.Author{
				DisplayName: displayName,
				Url:         authorUrl,
				Image:       &page.Image{Url: imageUrl},
			},
		},
	}

	// PageModel Database
	expectedErr := errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, expectedErr)
	actual, actualErr := logicService.Insert(insertReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorModel Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	authorRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, expectedErr)
	actual, actualErr = logicService.Insert(insertReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// ImageModel Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	authorRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	imageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, expectedErr)
	actual, actualErr = logicService.Insert(insertReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// PageNotExist
	expectedErr = errcode.Wrap(errcode.PageNotExist)
	pageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	authorRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	imageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogId, gomock.Any()).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Insert(insertReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// PageModel Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	authorRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	imageRepo.EXPECT().Insert(ctx, gomock.Any()).Return(nil, nil)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogId, gomock.Any()).Return(nil, expectedErr)
	actual, actualErr = logicService.Insert(insertReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
