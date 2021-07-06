package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Id        int       `json:id`
	Title     string    `json:title`
	Content   string    `json:content`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}

type ArticleDB interface {
	InsertArticle(db *gorm.DB, title string, content string)
	GetArticles(db *gorm.DB) []Article
}

type ArticleDao struct {
	articleDB ArticleDB
}

func NewArticleDB() ArticleDB {
	return &ArticleDao{}
}

func (a *ArticleDao) InsertArticle(db *gorm.DB, title string, content string) {
	art := Article{}
	art.Id = 0
	art.Title = title
	art.Content = content
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	art.CreatedAt = nowUTC.In(jst)
	fmt.Println(art.CreatedAt)
	db.Create(&art)
}

func (a *ArticleDao) GetArticles(db *gorm.DB) []Article {
	articles := []Article{}
	db.Find(&articles)
	return articles
}