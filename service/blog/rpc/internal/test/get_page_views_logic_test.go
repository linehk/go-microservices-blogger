package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetPageViews(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageViewsRepo := model.NewMockPageViewsModel(ctrl)
	logicService := logic.NewGetPageViewsLogic(ctx, &svc.ServiceContext{
		PageViewsModel: pageViewsRepo,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	var count int64 = 100
	getPageViews := &blog.GetPageViewsReq{BlogId: blogId}
	pageViewsModel := &model.PageViews{
		Id:       1,
		BlogUuid: blogId,
		Count:    sql.NullInt64{Int64: count, Valid: true},
	}

	expected := &blog.PageViews{
		Kind:   "blogger#page_views",
		BlogId: blogId,
		Counts: []*blog.Count{{
			TimeRange: "all",
			Count:     uint64(count),
		}},
	}

	// PageViewNotExist
	expectedErr := errcode.Wrap(errcode.PageViewNotExist)
	pageViewsRepo.EXPECT().FindOneByBlogUuid(ctx, blogId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.GetPageViews(getPageViews)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	pageViewsRepo.EXPECT().FindOneByBlogUuid(ctx, blogId).Return(nil, expectedErr)
	actual, actualErr = logicService.GetPageViews(getPageViews)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	pageViewsRepo.EXPECT().FindOneByBlogUuid(ctx, blogId).Return(pageViewsModel, nil)
	actual, actualErr = logicService.GetPageViews(getPageViews)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
