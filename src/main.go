package main

import (
	"github.com/KaoruOhbayashi/golang_echo/dao"
	"github.com/KaoruOhbayashi/golang_echo/tao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

//Connect mysql
func ConnectMysql() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "golang"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME   := "golang_echo"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true&loc=Asia%2FTokyo"
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func getIndexElement(a []dao.Article, e []string) interface{} {
	type Element struct {
		Art []dao.Article
		Err []string
	}
	m := Element{
		Art: a,
		Err: e,
	}
	return m
}

func getTemplate() *template.Template {
	// 関数の定義とマッピング
	funcMap := map[string]interface{}{
		"layout": func(At time.Time) string {
		layout := "2006年01月02日 15時04分"
		AtStr := At.Format(layout)
		return AtStr
		},
	}
	files := []string{"views/index.html","views/edit.html"}
	tname := filepath.Base(files[0])
	t := template.Must(template.New(tname).Funcs(funcMap).ParseFiles(files...))
	return t
}

func main() {

	//Connect mysql
	db := ConnectMysql()
	defer db.Close()

	e := echo.New()

	articleDB := dao.NewArticleDB() //generate article struct

	//Road templates
	t := &tao.Template{
		Templates: getTemplate(),
	}

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/index") // Redirect to home
	})

	//Output articles
	e.GET("index", func(c echo.Context) error{
		articles,_ := articleDB.GetArticles(db)
		err := []string{}
		element := getIndexElement(articles,err)
		return  c.Render(http.StatusOK, "index", element) // Redirect to home
	})

	//Post article
	e.POST("/insert", func(c echo.Context) error{
		err := articleDB.InsertArticle(db,c.FormValue("title"),c.FormValue("contents"))
		articles,_ := articleDB.GetArticles(db)
		element := getIndexElement(articles,err)
		return  c.Render(http.StatusOK, "index", element) // Redirect to home
	})

	e.POST("/delete", func(c echo.Context) error {
		articleDB.DeleteArticle(db,c.FormValue("deleteId"))
		return c.Redirect(http.StatusFound, "/index")
	})

	e.POST("/edit", func(c echo.Context) error {
		return c.Render(http.StatusOK,"edit",c.FormValue("editId"))
	})

	e.POST("/update", func(c echo.Context) error {
		err := articleDB.EditArticle(db,c.FormValue("editId"),c.FormValue("title"),c.FormValue("contents"))
		articles,_ := articleDB.GetArticles(db)
		element := getIndexElement(articles,err)
		return  c.Render(http.StatusOK, "index", element) // Redirect to home
	})

	e.Logger.Fatal(e.Start(":8080"))
}
