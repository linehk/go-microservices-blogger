package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/model"
	postmodel "github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	commentRepo := model.NewMockCommentModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		CommentModel: commentRepo,
		AuthorModel:  authorRepo,
		ImageModel:   imageRepo,
	})
	defer ctrl.Finish()

	blogUuid := uuid.NewString()
	postUuid := uuid.NewString()
	commentUuid := uuid.NewString()
	getReq := &comment.GetReq{
		BlogId:    blogUuid,
		CommentId: commentUuid,
		PostId:    postUuid,
	}

	status := "Status"
	published := time.Now()
	updated := time.Now()
	selfLink := "SelfLink"
	content := "Content"
	commentModel := &model.Comment{
		Id:        1,
		Uuid:      commentUuid,
		BlogUuid:  sql.NullString{String: blogUuid, Valid: true},
		PostUuid:  sql.NullString{String: postUuid, Valid: true},
		Status:    sql.NullString{String: status, Valid: true},
		Published: sql.NullTime{Time: published, Valid: true},
		Updated:   sql.NullTime{Time: updated, Valid: true},
		SelfLink:  sql.NullString{String: selfLink, Valid: true},
		Content:   sql.NullString{String: content, Valid: true},
	}

	authorUuid := uuid.NewString()
	displayName := "DisplayName"
	authorUrl := "Url"
	authorModel := &postmodel.Author{
		Id:          1,
		Uuid:        authorUuid,
		PostUuid:    "",
		PageUuid:    "",
		CommentUuid: commentUuid,
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

	expected := &comment.Comment{
		Kind:      "blogger#comment",
		Status:    status,
		Id:        commentUuid,
		InReplyTo: nil,
		Post:      &comment.Post{Id: postUuid},
		Blog:      &comment.Blog{Id: blogUuid},
		Published: timestamppb.New(published),
		Updated:   timestamppb.New(updated),
		SelfLink:  selfLink,
		Content:   content,
		Author: &comment.Author{
			Id:          authorUuid,
			DisplayName: displayName,
			Url:         authorUrl,
			Image:       &comment.Image{Url: imageUrl},
		},
	}

	// CommentNotExist
	expectedErr := errcode.Wrap(errcode.CommentNotExist)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorNotExist
	expectedErr = errcode.Wrap(errcode.AuthorNotExist)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// ImageNotExist
	expectedErr = errcode.Wrap(errcode.ImageNotExist)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(imageModel, nil)
	actual, actualErr = logicService.Get(getReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
