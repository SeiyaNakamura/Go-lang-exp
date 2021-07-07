package dao

import (
	mock_dao "github.com/KaoruOhbayashi/golang_echo/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestArticleDao_InsertArticle(t *testing.T) {
	type args struct {
		db      *gorm.DB
		title   string
		content string
	}
	tests := []struct {
		name   string
		prepareMockFn func(m *mock_dao.MockArticleDB)
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *mock_dao.MockArticleDB) {
				m.EXPECT().InsertArticle(gomock.Any(),"今日の朝食", "イチゴ").Return([]error{})
			},
			args: args{},
			want: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_dao.NewMockArticleDB(ctrl)

			tt.prepareMockFn(mock)

			a := &ArticleDao{
				articleDB: mock,
			}
			if got := a.InsertArticle(tt.args.db, tt.args.title, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDao.InsertArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDao_GetArticles(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		prepareMockFn func(m *mock_dao.MockArticleDB)
		args   args
		want   []Article
		want1  []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *mock_dao.MockArticleDB) {
				m.EXPECT().GetArticles(gomock.Any()).Return([]Article{},[]error{})
			},
			args: args{},
			want: []Article{},
			want1: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_dao.NewMockArticleDB(ctrl)

			tt.prepareMockFn(mock)

			a := &ArticleDao{
				articleDB: mock,
			}
			got, got1 := a.GetArticles(tt.args.db)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDao.GetArticles() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ArticleDao.GetArticles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArticleDao_DeleteArticle(t *testing.T) {
	type args struct {
		db *gorm.DB
		id string
	}
	tests := []struct {
		name   string
		prepareMockFn func(m *mock_dao.MockArticleDB)
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *mock_dao.MockArticleDB) {
				m.EXPECT().DeleteArticle(gomock.Any(),"1").Return([]error{})
			},
			args: args{},
			want: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_dao.NewMockArticleDB(ctrl)

			tt.prepareMockFn(mock)

			a := &ArticleDao{
				articleDB: mock,
			}
			if got := a.DeleteArticle(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDao.DeleteArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDao_EditArticle(t *testing.T) {
	type args struct {
		db       *gorm.DB
		id       string
		title    string
		contents string
	}
	tests := []struct {
		name   string
		prepareMockFn func(m *mock_dao.MockArticleDB)
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *mock_dao.MockArticleDB) {
				m.EXPECT().EditArticle(gomock.Any(),"1","フルスピード", "最高").Return([]error{})
			},
			args: args{},
			want: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_dao.NewMockArticleDB(ctrl)

			tt.prepareMockFn(mock)

			a := &ArticleDao{
				articleDB: mock,
			}
			if got := a.EditArticle(tt.args.db, tt.args.id, tt.args.title, tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDao.EditArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
