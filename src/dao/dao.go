package dao

import (
	"github.com/jinzhu/gorm"
	"time"
)

func GormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "golang"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME   := "golang_echo"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type article struct {
	Id       int       `json:id`
	Title      string    `json:title`
	Content    string    `json:content`
	CreatedAt time.Time `json:created_at`
	UpdatedAt  time.Time `json:updated_at`
}

func InsertArticle(db *gorm.DB, title string, content string) {
	a := article{}
	a.Id = 0
	a.Title = title
	a.Content = content
	a.CreatedAt = time.Now()
	db.Create(&a)
}

func GetArticles(db *gorm.DB) []article {
	a := []article{}
	db.Find(&a)
	return a
}