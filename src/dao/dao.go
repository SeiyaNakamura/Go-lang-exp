package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Article struct {
	Id        int       `json:id`
	Title     string    `json:"title" validate:"is_titleLen,is_onlySpace"`
	Content   string    `json:"content" validate:"is_contentLen,is_onlySpace"`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}

type ArticleDB interface {
	InsertArticle(db *gorm.DB, title string, content string) map[string]string
	GetArticles(db *gorm.DB) ([]Article,[]error)
	DeleteArticle(db *gorm.DB, id string) []error
	EditArticle(db *gorm.DB, id string, title string, content string) map[string]string
	GetEditArticle(db *gorm.DB,id string) Article
}

type ArticleDao struct {
	articleDB ArticleDB
}

func NewArticleDB() ArticleDB {
	return &ArticleDao{}
}

func (a *ArticleDao) InsertArticle(db *gorm.DB, title string, content string) map[string]string {
	art := Article{}
	art.Title = title
	art.Content = content

	//validation
	errors := ArticleValidation(&art)
	fmt.Println(errors)

	if len(errors) == 0 {
		db.Create(&art)
	}

	return errors
}

func (a *ArticleDao) GetArticles(db *gorm.DB) ([]Article,[]error) {
	articles := []Article{}
	result := db.Find(&articles)
	fmt.Println(result.GetErrors())
	return articles,result.GetErrors()
}

func (a *ArticleDao) DeleteArticle(db *gorm.DB, id string) []error {
	art := Article{}
	deleteId,_ := strconv.Atoi(id)
	result :=db.Delete(&art, deleteId)
	return result.GetErrors()
}

func (a *ArticleDao) EditArticle(db *gorm.DB, id string, title string, contents string) map[string]string{

	//get article what is edited
	art := Article{}
	editId ,_:= strconv.Atoi(id)
	db.Where("id = ?",editId).Select("created_at").Find(&art)

	art.Id = editId
	art.Title = title
	art.Content = contents
	art.UpdatedAt = time.Now()

	//validation
	errors := ArticleValidation(&art)

	if len(errors) == 0 {
		db.Save(&art)
	}

	return errors
}

func (a *ArticleDao) GetEditArticle(db *gorm.DB,id string) Article {
	art := Article{}
	editId ,_:= strconv.Atoi(id)
	db.Where("id = ?",editId).Find(&art)

	return art
}