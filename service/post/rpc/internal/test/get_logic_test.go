package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/comment"
	"github.com/linehk/go-microservices-blogger/service/comment/rpc/commentservice"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	imageRepo := model.NewMockImageModel(ctrl)
	authorRepo := model.NewMockAuthorModel(ctrl)
	commentService := commentservice.NewMockCommentService(ctrl)
	labelRepo := model.NewMockLabelModel(ctrl)
	locationRepo := model.NewMockLocationModel(ctrl)
	logicService := logic.NewGetLogic(ctx, &svc.ServiceContext{
		AuthorModel:    authorRepo,
		ImageModel:     imageRepo,
		LabelModel:     labelRepo,
		LocationModel:  locationRepo,
		PostModel:      postRepo,
		CommentService: commentService,
	})
	defer ctrl.Finish()

	postId := uuid.NewString()
	blogId := uuid.NewString()
	getReq := &post.GetReq{
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
	postModel := &model.Post{
		Id:             1,
		Uuid:           postId,
		BlogUuid:       blogId,
		Published:      sql.NullTime{Time: published, Valid: true},
		Updated:        sql.NullTime{Time: updated, Valid: true},
		Url:            sql.NullString{String: postUrl, Valid: true},
		SelfLink:       sql.NullString{String: postSelfLink, Valid: true},
		Title:          sql.NullString{String: postTitle, Valid: true},
		TitleLink:      sql.NullString{String: postTitleLink, Valid: true},
		Content:        sql.NullString{String: postContent, Valid: true},
		CustomMetaData: sql.NullString{String: customMetaData, Valid: true},
		Status:         sql.NullString{String: postStatus, Valid: true},
	}

	imageUuid1 := uuid.NewString()
	imageUuid2 := uuid.NewString()
	authorUuid := uuid.NewString()
	imageUrl1 := "imageUrl1"
	imageUrl2 := "imageUrl2"
	imageModelList := []*model.Image{
		{
			Id:         1,
			Uuid:       imageUuid1,
			PostUuid:   postId,
			AuthorUuid: authorUuid,
			Url:        sql.NullString{String: imageUrl1, Valid: true},
		}, {
			Id:         2,
			Uuid:       imageUuid2,
			PostUuid:   postId,
			AuthorUuid: authorUuid,
			Url:        sql.NullString{String: imageUrl2, Valid: true},
		},
	}

	pageUuid := uuid.NewString()
	commentUuid1 := uuid.NewString()
	commentUuid2 := uuid.NewString()
	displayName := "DisplayName"
	authorUrl := "authorUrl"
	authorModel := &model.Author{
		Id:          1,
		Uuid:        authorUuid,
		PostUuid:    postId,
		PageUuid:    pageUuid,
		CommentUuid: commentUuid1,
		DisplayName: sql.NullString{String: displayName, Valid: true},
		Url:         sql.NullString{String: authorUrl, Valid: true},
	}

	authorImageUuid := uuid.NewString()
	authorImageUrl := "authorImageUrl"
	authorImageModel := &model.Image{
		Id:         1,
		Uuid:       authorImageUuid,
		PostUuid:   postId,
		AuthorUuid: authorUuid,
		Url:        sql.NullString{String: authorImageUrl, Valid: true},
	}

	listCommentReq := &comment.ListReq{
		BlogId: blogId,
		PostId: postId,
	}

	commentStatus1 := "Status1"
	commentStatus2 := "Status2"
	inReplyToUuid := uuid.NewString()
	commentSelfLink1 := "commentSelfLink1"
	commentSelfLink2 := "commentSelfLink2"
	commentContent1 := "commentContent1"
	commentContent2 := "commentContent2"
	commentAuthorUuid := uuid.NewString()
	commentAuthorDisplayName := "commentAuthorDisplayName"
	commentAuthorUrl := "commentAuthorUrl"
	commentAuthorImageUrl := "commentAuthorImageUrl"
	listCommentResp := &comment.ListResp{
		Kind: "blogger#commentList",
		Items: []*comment.Comment{
			{
				Kind:      "blogger#comment",
				Status:    commentStatus1,
				Id:        commentUuid1,
				InReplyTo: &comment.InReplyTo{Id: inReplyToUuid},
				Post:      &comment.Post{Id: postId},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published),
				Updated:   timestamppb.New(updated),
				SelfLink:  commentSelfLink1,
				Content:   commentContent1,
				Author: &comment.Author{
					Id:          commentAuthorUuid,
					DisplayName: commentAuthorDisplayName,
					Url:         commentAuthorUrl,
					Image:       &comment.Image{Url: commentAuthorImageUrl},
				},
			}, {
				Kind:      "blogger#comment",
				Status:    commentStatus2,
				Id:        commentUuid2,
				InReplyTo: &comment.InReplyTo{Id: inReplyToUuid},
				Post:      &comment.Post{Id: postId},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published),
				Updated:   timestamppb.New(updated),
				SelfLink:  commentSelfLink2,
				Content:   commentContent2,
				Author: &comment.Author{
					Id:          commentAuthorUuid,
					DisplayName: commentAuthorDisplayName,
					Url:         commentAuthorUrl,
					Image:       &comment.Image{Url: commentAuthorImageUrl},
				},
			},
		},
	}

	labelUuid1 := uuid.NewString()
	labelUuid2 := uuid.NewString()
	labelValue1 := "labelValue1"
	labelValue2 := "labelValue2"
	labelModelList := []*model.Label{
		{
			Id:         1,
			Uuid:       labelUuid1,
			PostUuid:   postId,
			LabelValue: sql.NullString{String: labelValue1, Valid: true},
		}, {
			Id:         2,
			Uuid:       labelUuid2,
			PostUuid:   postId,
			LabelValue: sql.NullString{String: labelValue2, Valid: true},
		},
	}

	locationUuid := uuid.NewString()
	locationName := "locationName"
	locationLat := 1.1
	locationLng := 2.2
	locationSpan := "locationSpan"
	locationModel := &model.Location{
		Id:       1,
		Uuid:     locationUuid,
		PostUuid: postId,
		Name:     sql.NullString{String: locationName, Valid: true},
		Lat:      sql.NullFloat64{Float64: locationLat, Valid: true},
		Lng:      sql.NullFloat64{Float64: locationLng, Valid: true},
		Span:     sql.NullString{String: locationSpan, Valid: true},
	}

	expected := &post.Post{
		Kind:           "blogger#post",
		Id:             postId,
		Blog:           &post.Blog{Id: blogId},
		Published:      timestamppb.New(published),
		Updated:        timestamppb.New(updated),
		Url:            postUrl,
		SelfLink:       postSelfLink,
		Title:          postTitle,
		TitleLink:      postTitleLink,
		Content:        postContent,
		Images:         []*post.Image{{Url: imageUrl1}, {Url: imageUrl2}},
		CustomMetaData: customMetaData,
		Author: &post.Author{
			Id:          authorUuid,
			DisplayName: displayName,
			Url:         authorUrl,
			Image:       &post.Image{Url: authorImageUrl},
		},
		Replies: &post.Reply{
			TotalItems: 2,
			SelfLink:   "",
			Items: []*post.Comment{{
				Kind:      "blogger#comment",
				Status:    commentStatus1,
				Id:        commentUuid1,
				InReplyTo: &post.Comment_InReplyTo{Id: inReplyToUuid},
				Post:      &post.Comment_Post{Id: postId},
				Blog:      &post.Comment_Blog{Id: blogId},
				Published: timestamppb.New(published),
				Updated:   timestamppb.New(updated),
				SelfLink:  commentSelfLink1,
				Content:   commentContent1,
				Author: &post.Author{
					Id:          commentAuthorUuid,
					DisplayName: commentAuthorDisplayName,
					Url:         commentAuthorUrl,
					Image:       &post.Image{Url: commentAuthorImageUrl},
				},
			}, {
				Kind:      "blogger#comment",
				Status:    commentStatus2,
				Id:        commentUuid2,
				InReplyTo: &post.Comment_InReplyTo{Id: inReplyToUuid},
				Post:      &post.Comment_Post{Id: postId},
				Blog:      &post.Comment_Blog{Id: blogId},
				Published: timestamppb.New(published),
				Updated:   timestamppb.New(updated),
				SelfLink:  commentSelfLink2,
				Content:   commentContent2,
				Author: &post.Author{
					Id:          commentAuthorUuid,
					DisplayName: commentAuthorDisplayName,
					Url:         commentAuthorUrl,
					Image:       &post.Image{Url: commentAuthorImageUrl},
				},
			}},
		},
		Labels: []string{labelValue1, labelValue2},
		Location: &post.Location{
			Name: locationName,
			Lat:  float32(locationLat),
			Lng:  float32(locationLng),
			Span: locationSpan,
		},
		Status: postStatus,
	}

	// PostNotExist
	expectedErr := errcode.Wrap(errcode.PostNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// PostNotBelongToBlog
	expectedErr = errcode.Wrap(errcode.PostNotBelongToBlog)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	actual, actualErr = logicService.Get(&post.GetReq{BlogId: uuid.NewString(), PostId: postId})
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// ImageNotExist
	expectedErr = errcode.Wrap(errcode.ImageNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// AuthorNotExist
	expectedErr = errcode.Wrap(errcode.AuthorNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// author ImageNotExist
	expectedErr = errcode.Wrap(errcode.ImageNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Service
	expectedErr = errcode.Wrap(errcode.Service)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// LabelNotExist
	expectedErr = errcode.Wrap(errcode.LabelNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(listCommentResp, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(listCommentResp, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// LocationNotExist
	expectedErr = errcode.Wrap(errcode.LocationNotExist)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(listCommentResp, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId).Return(labelModelList, nil)
	locationRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(listCommentResp, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId).Return(labelModelList, nil)
	locationRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(nil, expectedErr)
	actual, actualErr = logicService.Get(getReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	postRepo.EXPECT().FindOneByUuid(ctx, postId).Return(postModel, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId).Return(imageModelList, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(authorModel, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid).Return(authorImageModel, nil)
	commentService.EXPECT().List(ctx, listCommentReq).Return(listCommentResp, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId).Return(labelModelList, nil)
	locationRepo.EXPECT().FindOneByPostUuid(ctx, postId).Return(locationModel, nil)
	actual, actualErr = logicService.Get(getReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
