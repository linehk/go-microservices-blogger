package test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/page/rpc/page"
	postmodel "github.com/linehk/go-microservices-blogger/service/post/rpc/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	pageRepo := model.NewMockPageModel(ctrl)
	authorRepo := postmodel.NewMockAuthorModel(ctrl)
	imageRepo := postmodel.NewMockImageModel(ctrl)
	logicService := logic.NewPatchLogic(ctx, &svc.ServiceContext{
		PageModel:   pageRepo,
		AuthorModel: authorRepo,
		ImageModel:  imageRepo,
	})
	defer ctrl.Finish()

	blogId := uuid.NewString()
	pageId := uuid.NewString()
	status := "Status"
	published := time.Now()
	updated := time.Now()
	pageUrl := "Url"
	selfLink := "SelfLink"
	title := "Title"
	content := "Content"

	displayName := "DisplayName"
	authorUrl := "Url"
	imageUrl := "imageUrl"
	patchReq := &page.PatchReq{
		BlogId: blogId,
		PageId: pageId,
		Page: &page.Page{
			Status:    status,
			Published: timestamppb.New(published),
			Updated:   timestamppb.New(updated),
			Url:       pageUrl,
			SelfLink:  selfLink,
			Title:     title,
			Content:   content,
			Author: &page.Author{
				DisplayName: displayName,
				Url:         authorUrl,
				Image:       &page.Image{Url: imageUrl},
			},
		},
	}

	// Database
	expectedErr := errcode.Wrap(errcode.Database)
	pageRepo.EXPECT().Update(ctx, gomock.Any()).Return(expectedErr)
	actual, actualErr := logicService.Patch(patchReq)
	assert.Nil(t, actual)
	assert.Equal(t, expectedErr, actualErr)
}
