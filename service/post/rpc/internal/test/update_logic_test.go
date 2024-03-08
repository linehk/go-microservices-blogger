package test

import (
	"context"
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

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	postRepo := model.NewMockPostModel(ctrl)
	imageRepo := model.NewMockImageModel(ctrl)
	authorRepo := model.NewMockAuthorModel(ctrl)
	commentService := commentservice.NewMockCommentService(ctrl)
	labelRepo := model.NewMockLabelModel(ctrl)
	locationRepo := model.NewMockLocationModel(ctrl)
	commentRepo := commentmodel.NewMockCommentModel(ctrl)
	logicService := logic.NewUpdateLogic(ctx, &svc.ServiceContext{
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
	published := time.Now()
	updated := time.Now()
	postUrl := "Url"
	postSelfLink := "postSelfLink"
	postTitle := "Title"
	postTitleLink := "postTitleLink"
	postContent := "Content"
	customMetaData := "CustomMetaData"
	postStatus := "Status"

	imageUrl1 := "imageUrl1"
	imageUrl2 := "imageUrl2"

	displayName := "DisplayName"
	authorUrl := "authorUrl"

	authorImageUrl := "authorImageUrl"

	commentStatus1 := "Status1"
	commentStatus2 := "Status2"
	commentSelfLink1 := "commentSelfLink1"
	commentSelfLink2 := "commentSelfLink2"
	commentContent1 := "commentContent1"
	commentContent2 := "commentContent2"
	commentAuthorDisplayName := "commentAuthorDisplayName"
	commentAuthorUrl := "commentAuthorUrl"
	commentAuthorImageUrl := "commentAuthorImageUrl"

	labelValue1 := "labelValue1"
	labelValue2 := "labelValue2"

	locationName := "locationName"
	locationLat := 1.1
	locationLng := 2.2
	locationSpan := "locationSpan"

	updateReq := &post.UpdateReq{
		BlogId: blogId,
		PostId: postId,
		Post: &post.Post{
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
				DisplayName: displayName,
				Url:         authorUrl,
				Image:       &post.Image{Url: authorImageUrl},
			},
			Replies: &post.Reply{
				Items: []*post.Comment{{
					Status:    commentStatus1,
					Published: timestamppb.New(published),
					Updated:   timestamppb.New(updated),
					SelfLink:  commentSelfLink1,
					Content:   commentContent1,
					Author: &post.Author{
						DisplayName: commentAuthorDisplayName,
						Url:         commentAuthorUrl,
						Image:       &post.Image{Url: commentAuthorImageUrl},
					},
				}, {
					Status:    commentStatus2,
					Published: timestamppb.New(published),
					Updated:   timestamppb.New(updated),
					SelfLink:  commentSelfLink2,
					Content:   commentContent2,
					Author: &post.Author{
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
		},
	}

	// Database
	expectedErr := errcode.Wrap(errcode.Database)
	postRepo.EXPECT().Update(ctx, gomock.Any()).Return(expectedErr)
	actual, actualErr := logicService.Update(updateReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
