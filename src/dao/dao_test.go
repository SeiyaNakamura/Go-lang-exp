package dao

import (
	mock_dao "github.com/KaoruOhbayashi/golang_echo/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestArticleDao_GetArticles(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		prepareMockFn func(m *mock_dao.MockArticleDB)
		args args
		want   []Article
	}{
		// TODO: Add test cases.
		{
			name: "success",
			prepareMockFn: func(m *mock_dao.MockArticleDB) {
				m.EXPECT().GetArticles(gomock.Any()).Return([]Article{})
			},
			args: args{},
			want: []Article{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_dao.NewMockArticleDB(ctrl)

			tt.prepareMockFn(mock)

			t1 := &ArticleDao{
				articleDB: mock,
			}

			if got := t1.GetArticles(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDao.GetArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
