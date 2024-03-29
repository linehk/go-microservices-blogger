// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/post/rpc/model (interfaces: AuthorModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_author_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model AuthorModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthorModel is a mock of AuthorModel interface.
type MockAuthorModel struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorModelMockRecorder
}

// MockAuthorModelMockRecorder is the mock recorder for MockAuthorModel.
type MockAuthorModelMockRecorder struct {
	mock *MockAuthorModel
}

// NewMockAuthorModel creates a new mock instance.
func NewMockAuthorModel(ctrl *gomock.Controller) *MockAuthorModel {
	mock := &MockAuthorModel{ctrl: ctrl}
	mock.recorder = &MockAuthorModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorModel) EXPECT() *MockAuthorModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAuthorModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockAuthorModel) FindOne(arg0 context.Context, arg1 int64) (*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockAuthorModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockAuthorModel)(nil).FindOne), arg0, arg1)
}

// FindOneByCommentUuid mocks base method.
func (m *MockAuthorModel) FindOneByCommentUuid(arg0 context.Context, arg1 string) (*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByCommentUuid", arg0, arg1)
	ret0, _ := ret[0].(*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByCommentUuid indicates an expected call of FindOneByCommentUuid.
func (mr *MockAuthorModelMockRecorder) FindOneByCommentUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByCommentUuid", reflect.TypeOf((*MockAuthorModel)(nil).FindOneByCommentUuid), arg0, arg1)
}

// FindOneByPageUuid mocks base method.
func (m *MockAuthorModel) FindOneByPageUuid(arg0 context.Context, arg1 string) (*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByPageUuid", arg0, arg1)
	ret0, _ := ret[0].(*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByPageUuid indicates an expected call of FindOneByPageUuid.
func (mr *MockAuthorModelMockRecorder) FindOneByPageUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByPageUuid", reflect.TypeOf((*MockAuthorModel)(nil).FindOneByPageUuid), arg0, arg1)
}

// FindOneByPostUuid mocks base method.
func (m *MockAuthorModel) FindOneByPostUuid(arg0 context.Context, arg1 string) (*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByPostUuid", arg0, arg1)
	ret0, _ := ret[0].(*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByPostUuid indicates an expected call of FindOneByPostUuid.
func (mr *MockAuthorModelMockRecorder) FindOneByPostUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByPostUuid", reflect.TypeOf((*MockAuthorModel)(nil).FindOneByPostUuid), arg0, arg1)
}

// FindOneByUuid mocks base method.
func (m *MockAuthorModel) FindOneByUuid(arg0 context.Context, arg1 string) (*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockAuthorModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockAuthorModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockAuthorModel) Insert(arg0 context.Context, arg1 *Author) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockAuthorModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAuthorModel)(nil).Insert), arg0, arg1)
}

// ListByPostUuid mocks base method.
func (m *MockAuthorModel) ListByPostUuid(arg0 context.Context, arg1 string) ([]*Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPostUuid", arg0, arg1)
	ret0, _ := ret[0].([]*Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPostUuid indicates an expected call of ListByPostUuid.
func (mr *MockAuthorModelMockRecorder) ListByPostUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPostUuid", reflect.TypeOf((*MockAuthorModel)(nil).ListByPostUuid), arg0, arg1)
}

// Update mocks base method.
func (m *MockAuthorModel) Update(arg0 context.Context, arg1 *Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuthorModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorModel)(nil).Update), arg0, arg1)
}
