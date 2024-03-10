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

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	commentRepo := model.NewMockCommentModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewListLogic(ctx, &svc.ServiceContext{
		CommentModel: commentRepo,
		AuthorModel:  authorRepo,
		ImageModel:   imageRepo,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	postId := uuid.NewString()
	listReq := &comment.ListReq{
		BlogId: blogId,
		PostId: postId,
	}

	commentUuid1 := uuid.NewString()
	status1 := "Status"
	published1 := time.Now()
	updated1 := time.Now()
	selfLink1 := "SelfLink"
	content1 := "Content"

	commentUuid2 := uuid.NewString()
	status2 := "Status"
	published2 := time.Now()
	updated2 := time.Now()
	selfLink2 := "SelfLink"
	content2 := "Content"
	commentModelList := []*model.Comment{
		{
			Id:        1,
			Uuid:      commentUuid1,
			BlogUuid:  sql.NullString{String: blogId, Valid: true},
			PostUuid:  sql.NullString{String: postId, Valid: true},
			Status:    sql.NullString{String: status1, Valid: true},
			Published: sql.NullTime{Time: published1, Valid: true},
			Updated:   sql.NullTime{Time: updated1, Valid: true},
			SelfLink:  sql.NullString{String: selfLink1, Valid: true},
			Content:   sql.NullString{String: content1, Valid: true},
		}, {
			Id:        2,
			Uuid:      commentUuid2,
			BlogUuid:  sql.NullString{String: blogId, Valid: true},
			PostUuid:  sql.NullString{String: postId, Valid: true},
			Status:    sql.NullString{String: status2, Valid: true},
			Published: sql.NullTime{Time: published2, Valid: true},
			Updated:   sql.NullTime{Time: updated2, Valid: true},
			SelfLink:  sql.NullString{String: selfLink2, Valid: true},
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
		PageUuid:    "",
		CommentUuid: commentUuid1,
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
		PageUuid:    "",
		CommentUuid: commentUuid2,
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

	expected := &comment.ListResp{
		Kind:          "blogger#commentList",
		NextPageToken: "",
		PrevPageToken: "",
		Items: []*comment.Comment{
			{
				Kind:      "blogger#comment",
				Status:    status1,
				Id:        commentUuid1,
				InReplyTo: nil,
				Post:      &comment.Post{Id: postId},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published1),
				Updated:   timestamppb.New(updated1),
				SelfLink:  selfLink1,
				Content:   content1,
				Author: &comment.Author{
					Id:          authorUuid1,
					DisplayName: displayName1,
					Url:         authorUrl1,
					Image:       &comment.Image{Url: imageUrl1},
				},
			}, {
				Kind:      "blogger#comment",
				Status:    status2,
				Id:        commentUuid2,
				InReplyTo: nil,
				Post:      &comment.Post{Id: postId},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published2),
				Updated:   timestamppb.New(updated2),
				SelfLink:  selfLink2,
				Content:   content2,
				Author: &comment.Author{
					Id:          authorUuid2,
					DisplayName: displayName2,
					Url:         authorUrl2,
					Image:       &comment.Image{Url: imageUrl2},
				},
			},
		},
	}

	// CommentNotExist
	expectedErr := errcode.Wrap(errcode.CommentNotExist)
	commentRepo.EXPECT().ListByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	commentRepo.EXPECT().ListByBlogUuidAndPostUuid(ctx, blogId, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorNotExist
	expectedErr = errcode.Wrap(errcode.AuthorNotExist)
	commentRepo.EXPECT().ListByBlogUuidAndPostUuid(ctx, blogId, postId).Return(commentModelList, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid1).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	commentRepo.EXPECT().ListByBlogUuidAndPostUuid(ctx, blogId, postId).Return(commentModelList, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid1).Return(authorModel1, nil)
	authorRepo.EXPECT().FindOneByCommentUuid(ctx, commentUuid2).Return(authorModel2, nil)

	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid1).Return(imageModel1, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid2).Return(imageModel2, nil)

	actual, actualErr = logicService.List(listReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
