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

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	imageRepo := model.NewMockImageModel(ctrl)
	authorRepo := model.NewMockAuthorModel(ctrl)
	commentService := commentservice.NewMockCommentService(ctrl)
	labelRepo := model.NewMockLabelModel(ctrl)
	locationRepo := model.NewMockLocationModel(ctrl)
	logicService := logic.NewListLogic(ctx, &svc.ServiceContext{
		AuthorModel:    authorRepo,
		ImageModel:     imageRepo,
		LabelModel:     labelRepo,
		LocationModel:  locationRepo,
		PostModel:      postRepo,
		CommentService: commentService,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	listReq := &post.ListReq{
		BlogId: blogId,
	}

	postId1 := uuid.NewString()
	published1 := time.Now()
	updated1 := time.Now()
	postUrl1 := "Url"
	postSelfLink1 := "postSelfLink1"
	postTitle1 := "Title"
	postTitleLink1 := "postTitleLink1"
	postContent1 := "Content"
	customMetaData1 := "CustomMetaData"
	postStatus1 := "Status"

	postId2 := uuid.NewString()
	published2 := time.Now()
	updated2 := time.Now()
	postUrl2 := "Url"
	postSelfLink2 := "postSelfLink1"
	postTitle2 := "Title"
	postTitleLink2 := "postTitleLink1"
	postContent2 := "Content"
	customMetaData2 := "CustomMetaData"
	postStatus2 := "Status"

	postModelList := []*model.Post{
		{
			Id:             1,
			Uuid:           postId1,
			BlogUuid:       sql.NullString{String: blogId, Valid: true},
			Published:      sql.NullTime{Time: published1, Valid: true},
			Updated:        sql.NullTime{Time: updated1, Valid: true},
			Url:            postUrl1,
			SelfLink:       sql.NullString{String: postSelfLink1, Valid: true},
			Title:          sql.NullString{String: postTitle1, Valid: true},
			TitleLink:      sql.NullString{String: postTitleLink1, Valid: true},
			Content:        sql.NullString{String: postContent1, Valid: true},
			CustomMetaData: sql.NullString{String: customMetaData1, Valid: true},
			Status:         sql.NullString{String: postStatus1, Valid: true},
		},
		{
			Id:             2,
			Uuid:           postId2,
			BlogUuid:       sql.NullString{String: blogId, Valid: true},
			Published:      sql.NullTime{Time: published2, Valid: true},
			Updated:        sql.NullTime{Time: updated2, Valid: true},
			Url:            postUrl2,
			SelfLink:       sql.NullString{String: postSelfLink2, Valid: true},
			Title:          sql.NullString{String: postTitle2, Valid: true},
			TitleLink:      sql.NullString{String: postTitleLink2, Valid: true},
			Content:        sql.NullString{String: postContent2, Valid: true},
			CustomMetaData: sql.NullString{String: customMetaData2, Valid: true},
			Status:         sql.NullString{String: postStatus2, Valid: true},
		},
	}

	imageUuid1 := uuid.NewString()
	imageUuid2 := uuid.NewString()
	authorUuid1 := uuid.NewString()
	imageUrl1 := "imageUrl1"
	imageUrl2 := "imageUrl2"
	imageModelList1 := []*model.Image{
		{
			Id:         1,
			Uuid:       imageUuid1,
			PostUuid:   sql.NullString{String: postId1, Valid: true},
			AuthorUuid: authorUuid1,
			Url:        sql.NullString{String: imageUrl1, Valid: true},
		}, {
			Id:         2,
			Uuid:       imageUuid2,
			PostUuid:   sql.NullString{String: postId1, Valid: true},
			AuthorUuid: authorUuid1,
			Url:        sql.NullString{String: imageUrl2, Valid: true},
		},
	}
	imageUuid3 := uuid.NewString()
	imageUuid4 := uuid.NewString()
	authorUuid2 := uuid.NewString()
	imageUrl3 := "imageUrl1"
	imageUrl4 := "imageUrl2"
	imageModelList2 := []*model.Image{
		{
			Id:         3,
			Uuid:       imageUuid3,
			PostUuid:   sql.NullString{String: postId2, Valid: true},
			AuthorUuid: authorUuid2,
			Url:        sql.NullString{String: imageUrl3, Valid: true},
		}, {
			Id:         4,
			Uuid:       imageUuid4,
			PostUuid:   sql.NullString{String: postId2, Valid: true},
			AuthorUuid: authorUuid2,
			Url:        sql.NullString{String: imageUrl4, Valid: true},
		},
	}

	displayName1 := "DisplayName"
	authorUrl1 := "authorUrl1"
	authorModel1 := &model.Author{
		Id:          1,
		Uuid:        authorUuid1,
		PostUuid:    postId1,
		PageUuid:    "",
		CommentUuid: "",
		DisplayName: sql.NullString{String: displayName1, Valid: true},
		Url:         sql.NullString{String: authorUrl1, Valid: true},
	}

	displayName2 := "DisplayName"
	authorUrl2 := "authorUrl1"
	authorModel2 := &model.Author{
		Id:          2,
		Uuid:        authorUuid2,
		PostUuid:    postId2,
		PageUuid:    "",
		CommentUuid: "",
		DisplayName: sql.NullString{String: displayName2, Valid: true},
		Url:         sql.NullString{String: authorUrl2, Valid: true},
	}

	authorImageUuid1 := uuid.NewString()
	authorImageUrl1 := "authorImageUrl1"
	authorImageModel1 := &model.Image{
		Id:         1,
		Uuid:       authorImageUuid1,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid1,
		Url:        sql.NullString{String: authorImageUrl1, Valid: true},
	}

	authorImageUuid2 := uuid.NewString()
	authorImageUrl2 := "authorImageUrl1"
	authorImageModel2 := &model.Image{
		Id:         2,
		Uuid:       authorImageUuid2,
		PostUuid:   sql.NullString{String: "", Valid: true},
		AuthorUuid: authorUuid2,
		Url:        sql.NullString{String: authorImageUrl2, Valid: true},
	}

	listCommentReq1 := &comment.ListReq{
		BlogId: blogId,
		PostId: postId1,
	}

	listCommentReq2 := &comment.ListReq{
		BlogId: blogId,
		PostId: postId2,
	}

	commentUuid1 := uuid.NewString()
	commentUuid2 := uuid.NewString()
	commentStatus1 := "Status1"
	commentStatus2 := "Status2"
	commentSelfLink1 := "commentSelfLink1"
	commentSelfLink2 := "commentSelfLink2"
	commentContent1 := "commentContent1"
	commentContent2 := "commentContent2"
	commentAuthorUuid1 := uuid.NewString()
	commentAuthorDisplayName1 := "commentAuthorDisplayName1"
	commentAuthorUrl1 := "commentAuthorUrl1"
	commentAuthorImageUrl1 := "commentAuthorImageUrl1"
	listCommentResp1 := &comment.ListResp{
		Kind: "blogger#commentList",
		Items: []*comment.Comment{
			{
				Kind:      "blogger#comment",
				Status:    commentStatus1,
				Id:        commentUuid1,
				InReplyTo: &comment.InReplyTo{Id: ""},
				Post:      &comment.Post{Id: postId1},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published1),
				Updated:   timestamppb.New(updated1),
				SelfLink:  commentSelfLink1,
				Content:   commentContent1,
				Author: &comment.Author{
					Id:          commentAuthorUuid1,
					DisplayName: commentAuthorDisplayName1,
					Url:         commentAuthorUrl1,
					Image:       &comment.Image{Url: commentAuthorImageUrl1},
				},
			}, {
				Kind:      "blogger#comment",
				Status:    commentStatus2,
				Id:        commentUuid2,
				InReplyTo: &comment.InReplyTo{Id: ""},
				Post:      &comment.Post{Id: postId1},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published1),
				Updated:   timestamppb.New(updated1),
				SelfLink:  commentSelfLink2,
				Content:   commentContent2,
				Author: &comment.Author{
					Id:          commentAuthorUuid1,
					DisplayName: commentAuthorDisplayName1,
					Url:         commentAuthorUrl1,
					Image:       &comment.Image{Url: commentAuthorImageUrl1},
				},
			},
		},
	}

	commentUuid3 := uuid.NewString()
	commentUuid4 := uuid.NewString()
	commentStatus3 := "Status1"
	commentStatus4 := "Status2"
	commentSelfLink3 := "commentSelfLink1"
	commentSelfLink4 := "commentSelfLink2"
	commentContent3 := "commentContent1"
	commentContent4 := "commentContent2"
	commentAuthorUuid2 := uuid.NewString()
	commentAuthorDisplayName2 := "commentAuthorDisplayName1"
	commentAuthorUrl2 := "commentAuthorUrl1"
	commentAuthorImageUrl2 := "commentAuthorImageUrl1"
	listCommentResp2 := &comment.ListResp{
		Kind: "blogger#commentList",
		Items: []*comment.Comment{
			{
				Kind:      "blogger#comment",
				Status:    commentStatus3,
				Id:        commentUuid3,
				InReplyTo: &comment.InReplyTo{Id: ""},
				Post:      &comment.Post{Id: postId2},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published2),
				Updated:   timestamppb.New(updated2),
				SelfLink:  commentSelfLink3,
				Content:   commentContent3,
				Author: &comment.Author{
					Id:          commentAuthorUuid2,
					DisplayName: commentAuthorDisplayName2,
					Url:         commentAuthorUrl2,
					Image:       &comment.Image{Url: commentAuthorImageUrl2},
				},
			}, {
				Kind:      "blogger#comment",
				Status:    commentStatus4,
				Id:        commentUuid4,
				InReplyTo: &comment.InReplyTo{Id: ""},
				Post:      &comment.Post{Id: postId2},
				Blog:      &comment.Blog{Id: blogId},
				Published: timestamppb.New(published2),
				Updated:   timestamppb.New(updated2),
				SelfLink:  commentSelfLink4,
				Content:   commentContent4,
				Author: &comment.Author{
					Id:          commentAuthorUuid2,
					DisplayName: commentAuthorDisplayName2,
					Url:         commentAuthorUrl2,
					Image:       &comment.Image{Url: commentAuthorImageUrl2},
				},
			},
		},
	}

	labelUuid1 := uuid.NewString()
	labelUuid2 := uuid.NewString()
	labelValue1 := "labelValue1"
	labelValue2 := "labelValue2"
	labelModelList1 := []*model.Label{
		{
			Id:         1,
			Uuid:       labelUuid1,
			PostUuid:   sql.NullString{String: postId1, Valid: true},
			LabelValue: sql.NullString{String: labelValue1, Valid: true},
		}, {
			Id:         2,
			Uuid:       labelUuid2,
			PostUuid:   sql.NullString{String: postId1, Valid: true},
			LabelValue: sql.NullString{String: labelValue2, Valid: true},
		},
	}

	labelUuid3 := uuid.NewString()
	labelUuid4 := uuid.NewString()
	labelValue3 := "labelValue1"
	labelValue4 := "labelValue2"
	labelModelList2 := []*model.Label{
		{
			Id:         3,
			Uuid:       labelUuid3,
			PostUuid:   sql.NullString{String: postId2, Valid: true},
			LabelValue: sql.NullString{String: labelValue3, Valid: true},
		}, {
			Id:         4,
			Uuid:       labelUuid4,
			PostUuid:   sql.NullString{String: postId2, Valid: true},
			LabelValue: sql.NullString{String: labelValue4, Valid: true},
		},
	}

	locationUuid1 := uuid.NewString()
	locationName1 := "locationName1"
	locationLat1 := 1.1
	locationLng1 := 2.2
	locationSpan1 := "locationSpan1"
	locationModel1 := &model.Location{
		Id:       1,
		Uuid:     locationUuid1,
		PostUuid: postId1,
		Name:     sql.NullString{String: locationName1, Valid: true},
		Lat:      sql.NullFloat64{Float64: locationLat1, Valid: true},
		Lng:      sql.NullFloat64{Float64: locationLng1, Valid: true},
		Span:     sql.NullString{String: locationSpan1, Valid: true},
	}

	locationUuid2 := uuid.NewString()
	locationName2 := "locationName1"
	locationLat2 := 1.1
	locationLng2 := 2.2
	locationSpan2 := "locationSpan1"
	locationModel2 := &model.Location{
		Id:       2,
		Uuid:     locationUuid2,
		PostUuid: postId2,
		Name:     sql.NullString{String: locationName2, Valid: true},
		Lat:      sql.NullFloat64{Float64: locationLat2, Valid: true},
		Lng:      sql.NullFloat64{Float64: locationLng2, Valid: true},
		Span:     sql.NullString{String: locationSpan2, Valid: true},
	}

	expected := &post.ListResp{
		Kind:          "blogger#postList",
		NextPageToken: "",
		Items: []*post.Post{
			{
				Kind:           "blogger#post",
				Id:             postId1,
				Blog:           &post.Blog{Id: blogId},
				Published:      timestamppb.New(published1),
				Updated:        timestamppb.New(updated1),
				Url:            postUrl1,
				SelfLink:       postSelfLink1,
				Title:          postTitle1,
				TitleLink:      postTitleLink1,
				Content:        postContent1,
				Images:         []*post.Image{{Url: imageUrl1}, {Url: imageUrl2}},
				CustomMetaData: customMetaData1,
				Author: &post.Author{
					Id:          authorUuid1,
					DisplayName: displayName1,
					Url:         authorUrl1,
					Image:       &post.Image{Url: authorImageUrl1},
				},
				Replies: &post.Reply{
					TotalItems: 2,
					SelfLink:   "",
					Items: []*post.Comment{{
						Kind:      "blogger#comment",
						Status:    commentStatus1,
						Id:        commentUuid1,
						InReplyTo: &post.Comment_InReplyTo{Id: ""},
						Post:      &post.Comment_Post{Id: postId1},
						Blog:      &post.Comment_Blog{Id: blogId},
						Published: timestamppb.New(published1),
						Updated:   timestamppb.New(updated1),
						SelfLink:  commentSelfLink1,
						Content:   commentContent1,
						Author: &post.Author{
							Id:          commentAuthorUuid1,
							DisplayName: commentAuthorDisplayName1,
							Url:         commentAuthorUrl1,
							Image:       &post.Image{Url: commentAuthorImageUrl1},
						},
					}, {
						Kind:      "blogger#comment",
						Status:    commentStatus2,
						Id:        commentUuid2,
						InReplyTo: &post.Comment_InReplyTo{Id: ""},
						Post:      &post.Comment_Post{Id: postId1},
						Blog:      &post.Comment_Blog{Id: blogId},
						Published: timestamppb.New(published1),
						Updated:   timestamppb.New(updated1),
						SelfLink:  commentSelfLink2,
						Content:   commentContent2,
						Author: &post.Author{
							Id:          commentAuthorUuid1,
							DisplayName: commentAuthorDisplayName1,
							Url:         commentAuthorUrl1,
							Image:       &post.Image{Url: commentAuthorImageUrl1},
						},
					}},
				},
				Labels: []string{labelValue1, labelValue2},
				Location: &post.Location{
					Name: locationName1,
					Lat:  float32(locationLat1),
					Lng:  float32(locationLng1),
					Span: locationSpan1,
				},
				Status: postStatus1,
			}, {
				Kind:           "blogger#post",
				Id:             postId2,
				Blog:           &post.Blog{Id: blogId},
				Published:      timestamppb.New(published2),
				Updated:        timestamppb.New(updated2),
				Url:            postUrl2,
				SelfLink:       postSelfLink2,
				Title:          postTitle2,
				TitleLink:      postTitleLink2,
				Content:        postContent2,
				Images:         []*post.Image{{Url: imageUrl3}, {Url: imageUrl4}},
				CustomMetaData: customMetaData2,
				Author: &post.Author{
					Id:          authorUuid2,
					DisplayName: displayName2,
					Url:         authorUrl2,
					Image:       &post.Image{Url: authorImageUrl2},
				},
				Replies: &post.Reply{
					TotalItems: 2,
					SelfLink:   "",
					Items: []*post.Comment{{
						Kind:      "blogger#comment",
						Status:    commentStatus3,
						Id:        commentUuid3,
						InReplyTo: &post.Comment_InReplyTo{Id: ""},
						Post:      &post.Comment_Post{Id: postId2},
						Blog:      &post.Comment_Blog{Id: blogId},
						Published: timestamppb.New(published2),
						Updated:   timestamppb.New(updated2),
						SelfLink:  commentSelfLink3,
						Content:   commentContent3,
						Author: &post.Author{
							Id:          commentAuthorUuid2,
							DisplayName: commentAuthorDisplayName2,
							Url:         commentAuthorUrl2,
							Image:       &post.Image{Url: commentAuthorImageUrl2},
						},
					}, {
						Kind:      "blogger#comment",
						Status:    commentStatus4,
						Id:        commentUuid4,
						InReplyTo: &post.Comment_InReplyTo{Id: ""},
						Post:      &post.Comment_Post{Id: postId2},
						Blog:      &post.Comment_Blog{Id: blogId},
						Published: timestamppb.New(published2),
						Updated:   timestamppb.New(updated2),
						SelfLink:  commentSelfLink4,
						Content:   commentContent4,
						Author: &post.Author{
							Id:          commentAuthorUuid2,
							DisplayName: commentAuthorDisplayName2,
							Url:         commentAuthorUrl2,
							Image:       &post.Image{Url: commentAuthorImageUrl2},
						},
					}},
				},
				Labels: []string{labelValue3, labelValue4},
				Location: &post.Location{
					Name: locationName2,
					Lat:  float32(locationLat2),
					Lng:  float32(locationLng2),
					Span: locationSpan2,
				},
				Status: postStatus2,
			},
		},
	}

	// PostNotExist
	expectedErr := errcode.Wrap(errcode.PostNotExist)
	postRepo.EXPECT().ListByBlogUuid(ctx, blogId).Return(nil, model.ErrNotFound)
	actual, actualErr := logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Database
	expectedErr = errcode.Wrap(errcode.Database)
	postRepo.EXPECT().ListByBlogUuid(ctx, blogId).Return(nil, expectedErr)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// ImageNotExist
	expectedErr = errcode.Wrap(errcode.ImageNotExist)
	postRepo.EXPECT().ListByBlogUuid(ctx, blogId).Return(postModelList, nil)
	imageRepo.EXPECT().ListByPostUuid(ctx, postId1).Return(nil, model.ErrNotFound)
	actual, actualErr = logicService.List(listReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)

	// Success
	postRepo.EXPECT().ListByBlogUuid(ctx, blogId).Return(postModelList, nil)

	imageRepo.EXPECT().ListByPostUuid(ctx, postId1).Return(imageModelList1, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId1).Return(authorModel1, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid1).Return(authorImageModel1, nil)
	commentService.EXPECT().List(ctx, listCommentReq1).Return(listCommentResp1, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId1).Return(labelModelList1, nil)
	locationRepo.EXPECT().FindOneByPostUuid(ctx, postId1).Return(locationModel1, nil)

	imageRepo.EXPECT().ListByPostUuid(ctx, postId2).Return(imageModelList2, nil)
	authorRepo.EXPECT().FindOneByPostUuid(ctx, postId2).Return(authorModel2, nil)
	imageRepo.EXPECT().FindOneByAuthorUuid(ctx, authorUuid2).Return(authorImageModel2, nil)
	commentService.EXPECT().List(ctx, listCommentReq2).Return(listCommentResp2, nil)
	labelRepo.EXPECT().ListByPostUuid(ctx, postId2).Return(labelModelList2, nil)
	locationRepo.EXPECT().FindOneByPostUuid(ctx, postId2).Return(locationModel2, nil)

	actual, actualErr = logicService.List(listReq)
	assert.Equal(t, expected, actual)
	assert.Nil(t, actualErr)
}
