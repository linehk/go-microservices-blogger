package test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/commentservice"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	imageRepo := model.NewMockImageModel(ctrl)
	authorRepo := model.NewMockAuthorModel(ctrl)
	commentService := commentservice.NewMockCommentService(ctrl)
	labelRepo := model.NewMockLabelModel(ctrl)
	locationRepo := model.NewMockLocationModel(ctrl)
	logicService := logic.NewSearchLogic(ctx, &svc.ServiceContext{
		AuthorModel:    authorRepo,
		ImageModel:     imageRepo,
		LabelModel:     labelRepo,
		LocationModel:  locationRepo,
		PostModel:      postRepo,
		CommentService: commentService,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	title := "title"
	searchReq := &post.SearchReq{
		BlogId: blogId,
		Q:      title,
	}

	// PostNotExist
	expectedErr := errcode.Wrap(errcode.PostNotExist)
	postRepo.EXPECT().SearchByTitle(ctx, blogId, title).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Search(searchReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().SearchByTitle(ctx, blogId, title).Return(nil, expectedErr)
	actual, actualErr = logicService.Search(searchReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
