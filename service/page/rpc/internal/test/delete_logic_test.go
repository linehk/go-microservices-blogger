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
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageRepo := model.NewMockPageModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewDeleteLogic(ctx, &svc.ServiceContext{
		PageModel:   pageRepo,
		AuthorModel: authorRepo,
		ImageModel:  imageRepo,
	})
	defer ctrl.Finish()

	blogUuid := uuid.NewString()
	pageUuid := uuid.NewString()
	deleteReq := &page.DeleteReq{
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
	var pagePrimaryKey int64 = 1
	pageModel := &model.Page{
		Id:        pagePrimaryKey,
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

	expected := &page.EmptyResp{}
	
	// PageNotExist
	expectedErr := errcode.Wrap(errcode.PageNotExist)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	pageRepo.EXPECT().Delete(ctx, pagePrimaryKey).Return(expectedErr)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	pageRepo.EXPECT().FindOneByBlogUuidAndPageUuid(ctx, blogUuid, pageUuid).Return(pageModel, nil)
	pageRepo.EXPECT().Delete(ctx, pagePrimaryKey).Return(nil)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
