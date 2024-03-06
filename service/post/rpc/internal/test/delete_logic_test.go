package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	logicService := logic.NewDeleteLogic(ctx, &svc.ServiceContext{
		PostModel: postRepo,
	})
	defer ctrl.Finish()

	postId := uuid.NewString()
	blogId := uuid.NewString()
	deleteReq := &post.DeleteReq{
		BlogId: blogId,
		PostId: postId,
	}

	published := time.Now()
	updated := time.Now()
	postUrl := "Url"
	postSelfLink := "postSelfLink"
	postTitle := "Title"
	postTitleLink := "postTitleLink"
	postContent := "Content"
	customMetaData := "CustomMetaData"
	postStatus := "Status"
	var postPrimaryKey int64 = 1
	postModel := &model.Post{
		Id:             postPrimaryKey,
		Uuid:           postId,
		BlogUuid:       blogId,
		Published:      sql.NullTime{Time: published, Valid: true},
		Updated:        sql.NullTime{Time: updated, Valid: true},
		Url:            postUrl,
		SelfLink:       sql.NullString{String: postSelfLink, Valid: true},
		Title:          sql.NullString{String: postTitle, Valid: true},
		TitleLink:      sql.NullString{String: postTitleLink, Valid: true},
		Content:        sql.NullString{String: postContent, Valid: true},
		CustomMetaData: sql.NullString{String: customMetaData, Valid: true},
		Status:         sql.NullString{String: postStatus, Valid: true},
	}

	expected := &post.EmptyResp{}

	// PostNotExist
	expectedErr := errcode.Wrap(errcode.PostNotExist)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(postModel, nil)
	postRepo.EXPECT().Delete(ctx, postPrimaryKey).Return(expectedErr)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(postModel, nil)
	postRepo.EXPECT().Delete(ctx, postPrimaryKey).Return(nil)
	actual, actualErr = logicService.Delete(deleteReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
