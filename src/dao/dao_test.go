package dao

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getVirtualDB(a Article) *gorm.DB {
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	db = db.DropTable(a)
	db = db.AutoMigrate(a)
	return db
}

func TestArticleDao_InsertArticle(t *testing.T) {
	type args struct {
		db      *gorm.DB
		title   string
		content string
	}
	tests := []struct {
		name   string
		args   args
		prepareMockFn func(m *MockArticleDB)
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "Title_Error",
			args: args{db: getVirtualDB(Article{}),title: "", content: "イチゴ"},
			prepareMockFn: func(m *MockArticleDB) {},
			want: []string {"error message for Title"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockArticleDB(ctrl)

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
		prepareMockFn func(m *MockArticleDB)
		args   args
		want   []Article
		want1  []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *MockArticleDB) {
				//m.EXPECT().GetArticles(getVirtualDB(Article{})).Return([]Article{},[]error{})
			},
			args: args{db: getVirtualDB(Article{})},
			want: []Article{},
			want1: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockArticleDB(ctrl)

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
		prepareMockFn func(m *MockArticleDB)
		args   args
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *MockArticleDB) {
				//m.EXPECT().DeleteArticle(gomock.Any(),gomock.Any()).Return([]error{})
			},
			args: args{db: getVirtualDB(Article{}),id: "1"},
			want: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockArticleDB(ctrl)

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
		prepareMockFn func(m *MockArticleDB)
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *MockArticleDB) {
				//m.EXPECT().EditArticle(getVirtualDB(),"1","フルスピード", "最高").Return([]string{})
			},
			args: args{db: getVirtualDB(Article{}),id: "1",title: "", contents: "バナナ"},
			want: []string{"error message for Title"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockArticleDB(ctrl)

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
