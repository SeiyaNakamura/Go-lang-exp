package main

import (
	"fmt"
	"github.com/KaoruOhbayashi/golang_echo/dao"
	"github.com/KaoruOhbayashi/golang_echo/tao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

//Connect mysql
func ConnectMysql() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "golang"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME   := "golang_echo"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	//Connect mysql
	db := ConnectMysql()
	defer db.Close()

	e := echo.New()

	articleDB := dao.NewArticleDB() //generate article struct

	//Road templates
	t := &tao.Template{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/index") // Redirect to home
	})

	//Output articles
	e.GET("index", func(c echo.Context) error{
		fmt.Println("hello")
		articles := articleDB.GetArticles(db)
		fmt.Println(articles[0].UpdatedAt)
		return  c.Render(http.StatusOK, "index", articles)
	})

	//Post article
	e.POST("/insert", func(c echo.Context) error{
		articleDB.InsertArticle(db,c.FormValue("title"),c.FormValue("contents"))
		return c.Redirect(http.StatusFound, "/index") // Redirect to home
	})

	e.Logger.Fatal(e.Start(":8080"))
}
