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
)

func TestApprove(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	commentRepo := model.NewMockCommentModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewApproveLogic(ctx, &svc.ServiceContext{
		CommentModel: commentRepo,
		AuthorModel:  authorRepo,
		ImageModel:   imageRepo,
	})
	defer ctrl.Finish()

	blogUuid := uuid.NewString()
	postUuid := uuid.NewString()
	commentUuid := uuid.NewString()
	approveReq := &comment.ApproveReq{
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

	// CommentNotExist
	expectedErr := errcode.Wrap(errcode.CommentNotExist)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Approve(approveReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Approve(approveReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().FindOneByUuid(ctx, commentUuid).Return(commentModel, nil)
	commentRepo.EXPECT().Update(ctx, gomock.Any()).Return(expectedErr)
	actual, actualErr = logicService.Approve(approveReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
