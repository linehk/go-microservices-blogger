// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"github.com/linehk/go-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-blogger/service/blog/rpc/internal/logic"
	"github.com/linehk/go-blogger/service/blog/rpc/internal/svc"
)

type BlogServiceServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedBlogServiceServer
}

func NewBlogServiceServer(svcCtx *svc.ServiceContext) *BlogServiceServer {
	return &BlogServiceServer{
		svcCtx: svcCtx,
	}
}
