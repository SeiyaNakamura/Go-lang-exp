// Code generated by MockGen. DO NOT EDIT.
// Source: dao/dao.go

// Package mock_dao is a generated GoMock package.
package dao

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockArticleDB is a mock of ArticleDB interface.
type MockArticleDB struct {
	ctrl     *gomock.Controller
	recorder *MockArticleDBMockRecorder
}

// MockArticleDBMockRecorder is the mock recorder for MockArticleDB.
type MockArticleDBMockRecorder struct {
	mock *MockArticleDB
}

// NewMockArticleDB creates a new mock instance.
func NewMockArticleDB(ctrl *gomock.Controller) *MockArticleDB {
	mock := &MockArticleDB{ctrl: ctrl}
	mock.recorder = &MockArticleDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleDB) EXPECT() *MockArticleDBMockRecorder {
	return m.recorder
}

// DeleteArticle mocks base method.
func (m *MockArticleDB) DeleteArticle(db *gorm.DB, id string) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", db, id)
	ret0, _ := ret[0].([]error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockArticleDBMockRecorder) DeleteArticle(db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockArticleDB)(nil).DeleteArticle), db, id)
}

// EditArticle mocks base method.
func (m *MockArticleDB) EditArticle(db *gorm.DB, id, title, content string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditArticle", db, id, title, content)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// EditArticle indicates an expected call of EditArticle.
func (mr *MockArticleDBMockRecorder) EditArticle(db, id, title, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditArticle", reflect.TypeOf((*MockArticleDB)(nil).EditArticle), db, id, title, content)
}

// GetArticles mocks base method.
func (m *MockArticleDB) GetArticles(db *gorm.DB) ([]Article, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", db)
	ret0, _ := ret[0].([]Article)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockArticleDBMockRecorder) GetArticles(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockArticleDB)(nil).GetArticles), db)
}

// GetEditArticle mocks base method.
func (m *MockArticleDB) GetEditArticle(db *gorm.DB, id string) Article {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEditArticle", db, id)
	ret0, _ := ret[0].(Article)
	return ret0
}

// GetEditArticle indicates an expected call of GetEditArticle.
func (mr *MockArticleDBMockRecorder) GetEditArticle(db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEditArticle", reflect.TypeOf((*MockArticleDB)(nil).GetEditArticle), db, id)
}

// InsertArticle mocks base method.
func (m *MockArticleDB) InsertArticle(db *gorm.DB, title, content string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertArticle", db, title, content)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockArticleDBMockRecorder) InsertArticle(db, title, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockArticleDB)(nil).InsertArticle), db, title, content)
}
