// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package server

import (
	"context"

	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/logic"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/post/rpc/post"
)

type PostServiceServer struct {
	svcCtx *svc.ServiceContext
	post.UnimplementedPostServiceServer
}

func NewPostServiceServer(svcCtx *svc.ServiceContext) *PostServiceServer {
	return &PostServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *PostServiceServer) List(ctx context.Context, in *post.ListReq) (*post.ListResp, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *PostServiceServer) Get(ctx context.Context, in *post.GetReq) (*post.Post, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *PostServiceServer) Search(ctx context.Context, in *post.SearchReq) (*post.SearchResp, error) {
	l := logic.NewSearchLogic(ctx, s.svcCtx)
	return l.Search(in)
}

func (s *PostServiceServer) Insert(ctx context.Context, in *post.InsertReq) (*post.Post, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *PostServiceServer) Delete(ctx context.Context, in *post.DeleteReq) (*post.EmptyResp, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *PostServiceServer) GetByPath(ctx context.Context, in *post.GetByPathReq) (*post.Post, error) {
	l := logic.NewGetByPathLogic(ctx, s.svcCtx)
	return l.GetByPath(in)
}

func (s *PostServiceServer) Patch(ctx context.Context, in *post.PatchReq) (*post.Post, error) {
	l := logic.NewPatchLogic(ctx, s.svcCtx)
	return l.Patch(in)
}

func (s *PostServiceServer) Update(ctx context.Context, in *post.UpdateReq) (*post.Post, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *PostServiceServer) Publish(ctx context.Context, in *post.PublishReq) (*post.Post, error) {
	l := logic.NewPublishLogic(ctx, s.svcCtx)
	return l.Publish(in)
}

func (s *PostServiceServer) Revert(ctx context.Context, in *post.RevertReq) (*post.Post, error) {
	l := logic.NewRevertLogic(ctx, s.svcCtx)
	return l.Revert(in)
}

func (s *PostServiceServer) GetPostUserInfos(ctx context.Context, in *post.GetPostUserInfosReq) (*post.PostUserInfos, error) {
	l := logic.NewGetPostUserInfosLogic(ctx, s.svcCtx)
	return l.GetPostUserInfos(in)
}

func (s *PostServiceServer) ListPostUserInfos(ctx context.Context, in *post.ListPostUserInfosReq) (*post.ListPostUserInfosResp, error) {
	l := logic.NewListPostUserInfosLogic(ctx, s.svcCtx)
	return l.ListPostUserInfos(in)
}
