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

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageRepo := model.NewMockPageModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewListLogic(ctx, &svc.ServiceContext{
		PageModel:   pageRepo,
		AuthorModel: authorRepo,
		ImageModel:  imageRepo,
	})
	defer ctrl.Finish()

	blogUuid := uuid.NewString()
	listReq := &page.ListReq{
		BlogId: blogUuid,
	}

	pageUuid1 := uuid.NewString()
	status1 := "Status"
	published1 := time.Now()
	updated1 := time.Now()
	pageUrl1 := "Url"
	selfLink1 := "SelfLink"
	title1 := "Title"
	content1 := "Content"

	pageUuid2 := uuid.NewString()
	status2 := "Status"
	published2 := time.Now()
	updated2 := time.Now()
	pageUrl2 := "Url"
	selfLink2 := "SelfLink"
	title2 := "Title"
	content2 := "Content"

	pageModelList := []*model.Page{
		{
			Id:        1,
			Uuid:      pageUuid1,
			BlogUuid:  sql.NullString{String: blogUuid, Valid: true},
			Status:    sql.NullString{String: status1, Valid: true},
			Published: sql.NullTime{Time: published1, Valid: true},
			Updated:   sql.NullTime{Time: updated1, Valid: true},
			Url:       sql.NullString{String: pageUrl1, Valid: true},
			SelfLink:  sql.NullString{String: selfLink1, Valid: true},
			Title:     sql.NullString{String: title1, Valid: true},
			Content:   sql.NullString{String: content1, Valid: true},
		}, {
			Id:        2,
			Uuid:      pageUuid2,
			BlogUuid:  sql.NullString{String: blogUuid, Valid: true},
			Status:    sql.NullString{String: status2, Valid: true},
			Published: sql.NullTime{Time: published2, Valid: true},
			Updated:   sql.NullTime{Time: updated2, Valid: true},
			Url:       sql.NullString{String: pageUrl2, Valid: true},
			SelfLink:  sql.NullString{String: selfLink2, Valid: true},
			Title:     sql.NullString{String: title2, Valid: true},
			Content:   sql.NullString{String: content2, Valid: true},
		},
	}

	authorUuid1 := uuid.NewString()
	displayName1 := "DisplayName"
	authorUrl1 := "Url"
	authorModel1 := &postmodel.Author{
		Id:          1,
		Uuid:        authorUuid1,
		PostUuid:    "",
		PageUuid:    pageUuid1,
		CommentUuid: "",
		DisplayName: sql.NullString{String: displayName1, Valid: true},
		Url:         sql.NullString{String: authorUrl1, Valid: true},
	}

	authorUuid2 := uuid.NewString()
	displayName2 := "DisplayName"
	authorUrl2 := "Url"
	authorModel2 := &postmodel.Author{
		Id:          2,
		Uuid:        authorUuid2,
		PostUuid:    "",
		PageUuid:    pageUuid2,
		CommentUuid: "",
		DisplayName: sql.NullString{String: displayName2, Valid: true},
		Url:         sql.NullString{String: authorUrl2, Valid: true},
	}

	imageUuid1 := uuid.NewString()
	imageUrl1 := "imageUrl"
	imageModel1 := &postmodel.Image{
		Id:         1,
		Uuid:       imageUuid1,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid1,
		Url:        sql.NullString{String: imageUrl1, Valid: true},
	}

	imageUuid2 := uuid.NewString()
	imageUrl2 := "imageUrl"
	imageModel2 := &postmodel.Image{
		Id:         2,
		Uuid:       imageUuid2,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid2,
		Url:        sql.NullString{String: imageUrl2, Valid: true},
	}

	expected := &page.ListResp{
		Kind: "blogger#pageList",
		Items: []*page.Page{
			{
				Kind:      "blogger#page",
				Id:        pageUuid1,
				Status:    status1,
				Blog:      &page.Blog{Id: blogUuid},
				Published: timestamppb.New(published1),
				Updated:   timestamppb.New(updated1),
				Url:       pageUrl1,
				SelfLink:  selfLink1,
				Title:     title1,
				Content:   content1,
				Author: &page.Author{
					Id:          authorUuid1,
					DisplayName: displayName1,
					Url:         authorUrl1,
					Image:       &page.Image{Url: imageUrl1},
				},
			}, {
				Kind:      "blogger#page",
				Id:        pageUuid2,
				Status:    status2,
				Blog:      &page.Blog{Id: blogUuid},
				Published: timestamppb.New(published2),
				Updated:   timestamppb.New(updated2),
				Url:       pageUrl2,
				SelfLink:  selfLink2,
				Title:     title2,
				Content:   content2,
				Author: &page.Author{
					Id:          authorUuid2,
					DisplayName: displayName2,
					Url:         authorUrl2,
					Image:       &page.Image{Url: imageUrl2},
				},
			},
		},
	}

	// PageNotExist
	expectedErr := errcode.Wrap(errcode.PageNotExist)
	pageRepo.EXPECT().ListByBlogUuid(ctx, blogUuid).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().ListByBlogUuid(ctx, blogUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorNotExist
	expectedErr = errcode.Wrap(errcode.AuthorNotExist)
	pageRepo.EXPECT().ListByBlogUuid(ctx, blogUuid).Return(pageModelList, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid1).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	pageRepo.EXPECT().ListByBlogUuid(ctx, blogUuid).Return(pageModelList, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid1).Return(authorModel1, nil)
	authorRepo.EXPECT().FindOneByPageUuid(ctx, pageUuid2).Return(authorModel2, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid1).Return(imageModel1, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid2).Return(imageModel2, nil)

	actual, actualErr = logicService.List(listReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
