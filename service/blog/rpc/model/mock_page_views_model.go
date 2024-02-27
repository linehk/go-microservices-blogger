// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/blog/rpc/model (interfaces: PageViewsModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_page_views_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/blog/rpc/model github.com/linehk/go-microservices-blogger/service/blog/rpc/model PageViewsModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPageViewsModel is a mock of PageViewsModel interface.
type MockPageViewsModel struct {
	ctrl     *gomock.Controller
	recorder *MockPageViewsModelMockRecorder
}

// MockPageViewsModelMockRecorder is the mock recorder for MockPageViewsModel.
type MockPageViewsModelMockRecorder struct {
	mock *MockPageViewsModel
}

// NewMockPageViewsModel creates a new mock instance.
func NewMockPageViewsModel(ctrl *gomock.Controller) *MockPageViewsModel {
	mock := &MockPageViewsModel{ctrl: ctrl}
	mock.recorder = &MockPageViewsModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageViewsModel) EXPECT() *MockPageViewsModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPageViewsModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPageViewsModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPageViewsModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockPageViewsModel) FindOne(arg0 context.Context, arg1 int64) (*PageViews, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*PageViews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPageViewsModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPageViewsModel)(nil).FindOne), arg0, arg1)
}

// FindOneByBlogUuid mocks base method.
func (m *MockPageViewsModel) FindOneByBlogUuid(arg0 context.Context, arg1 string) (*PageViews, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByBlogUuid", arg0, arg1)
	ret0, _ := ret[0].(*PageViews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByBlogUuid indicates an expected call of FindOneByBlogUuid.
func (mr *MockPageViewsModelMockRecorder) FindOneByBlogUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByBlogUuid", reflect.TypeOf((*MockPageViewsModel)(nil).FindOneByBlogUuid), arg0, arg1)
}

// FindOneByUuid mocks base method.
func (m *MockPageViewsModel) FindOneByUuid(arg0 context.Context, arg1 string) (*PageViews, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*PageViews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockPageViewsModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockPageViewsModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockPageViewsModel) Insert(arg0 context.Context, arg1 *PageViews) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPageViewsModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPageViewsModel)(nil).Insert), arg0, arg1)
}

// Update mocks base method.
func (m *MockPageViewsModel) Update(arg0 context.Context, arg1 *PageViews) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPageViewsModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPageViewsModel)(nil).Update), arg0, arg1)
}