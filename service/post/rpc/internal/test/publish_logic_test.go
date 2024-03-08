package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/commentservice"
	commentmodel "github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	imageRepo := model.NewMockImageModel(ctrl)
	authorRepo := model.NewMockAuthorModel(ctrl)
	commentService := commentservice.NewMockCommentService(ctrl)
	labelRepo := model.NewMockLabelModel(ctrl)
	locationRepo := model.NewMockLocationModel(ctrl)
	commentRepo := commentmodel.NewMockCommentModel(ctrl)
	logicService := logic.NewPublishLogic(ctx, &svc.ServiceContext{
		AuthorModel:    authorRepo,
		ImageModel:     imageRepo,
		LabelModel:     labelRepo,
		LocationModel:  locationRepo,
		PostModel:      postRepo,
		CommentService: commentService,
		CommentModel:   commentRepo,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	postId := uuid.NewString()
	publishDate := timestamppb.Now()
	publishReq := &post.PublishReq{
		BlogId:      blogId,
		PostId:      postId,
		PublishDate: publishDate,
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
	postModel := &model.Post{
		Id:             1,
		Uuid:           postId,
		BlogUuid:       sql.NullString{String: blogId, Valid: true},
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

	// PostNotExist
	expectedErr := errcode.Wrap(errcode.PostNotExist)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Publish(publishReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// FindOneByBlogUuidAndPostUuid Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Publish(publishReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Update Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByBlogUuidAndPostUuid(ctx, blogId, postId).Return(postModel, nil)
	postRepo.EXPECT().Update(ctx, gomock.Any()).Return(expectedErr)
	actual, actualErr = logicService.Publish(publishReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
