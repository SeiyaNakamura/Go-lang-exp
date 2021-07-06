// Code generated by MockGen. DO NOT EDIT.
// Source: dao/dao.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	reflect "reflect"

	dao "github.com/KaoruOhbayashi/golang_echo/dao"
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

// GetArticles mocks base method.
func (m *MockArticleDB) GetArticles(db *gorm.DB) []dao.Article {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", db)
	ret0, _ := ret[0].([]dao.Article)
	return ret0
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockArticleDBMockRecorder) GetArticles(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockArticleDB)(nil).GetArticles), db)
}

// InsertArticle mocks base method.
func (m *MockArticleDB) InsertArticle(db *gorm.DB, title, content string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertArticle", db, title, content)
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockArticleDBMockRecorder) InsertArticle(db, title, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockArticleDB)(nil).InsertArticle), db, title, content)
}
