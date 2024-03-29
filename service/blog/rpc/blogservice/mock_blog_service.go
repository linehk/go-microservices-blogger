// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/blog/rpc/blogservice/blog_service.go
//
// Generated by this command:
//
//	mockgen -source=./service/blog/rpc/blogservice/blog_service.go -destination=./service/blog/rpc/blogservice/mock_blog_service.go -package=blogservice -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/blogservice github.com/linehk/go-microservices-blogger/service/blog/rpc/blogservice
//

// Package blogservice is a generated GoMock package.
package blogservice

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBlogService is a mock of BlogService interface.
type MockBlogService struct {
	ctrl     *gomock.Controller
	recorder *MockBlogServiceMockRecorder
}

// MockBlogServiceMockRecorder is the mock recorder for MockBlogService.
type MockBlogServiceMockRecorder struct {
	mock *MockBlogService
}

// NewMockBlogService creates a new mock instance.
func NewMockBlogService(ctrl *gomock.Controller) *MockBlogService {
	mock := &MockBlogService{ctrl: ctrl}
	mock.recorder = &MockBlogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogService) EXPECT() *MockBlogServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBlogService) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Blog, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlogServiceMockRecorder) Get(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlogService)(nil).Get), varargs...)
}

// GetBlogUserInfos mocks base method.
func (m *MockBlogService) GetBlogUserInfos(ctx context.Context, in *GetBlogUserInfosReq, opts ...grpc.CallOption) (*BlogUserInfos, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlogUserInfos", varargs...)
	ret0, _ := ret[0].(*BlogUserInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlogUserInfos indicates an expected call of GetBlogUserInfos.
func (mr *MockBlogServiceMockRecorder) GetBlogUserInfos(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogUserInfos", reflect.TypeOf((*MockBlogService)(nil).GetBlogUserInfos), varargs...)
}

// GetByUrl mocks base method.
func (m *MockBlogService) GetByUrl(ctx context.Context, in *GetByUrlReq, opts ...grpc.CallOption) (*Blog, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByUrl", varargs...)
	ret0, _ := ret[0].(*Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUrl indicates an expected call of GetByUrl.
func (mr *MockBlogServiceMockRecorder) GetByUrl(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUrl", reflect.TypeOf((*MockBlogService)(nil).GetByUrl), varargs...)
}

// GetPageViews mocks base method.
func (m *MockBlogService) GetPageViews(ctx context.Context, in *GetPageViewsReq, opts ...grpc.CallOption) (*PageViews, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPageViews", varargs...)
	ret0, _ := ret[0].(*PageViews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPageViews indicates an expected call of GetPageViews.
func (mr *MockBlogServiceMockRecorder) GetPageViews(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPageViews", reflect.TypeOf((*MockBlogService)(nil).GetPageViews), varargs...)
}

// ListByUser mocks base method.
func (m *MockBlogService) ListByUser(ctx context.Context, in *ListByUserReq, opts ...grpc.CallOption) (*ListByUserResp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByUser", varargs...)
	ret0, _ := ret[0].(*ListByUserResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUser indicates an expected call of ListByUser.
func (mr *MockBlogServiceMockRecorder) ListByUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUser", reflect.TypeOf((*MockBlogService)(nil).ListByUser), varargs...)
}
