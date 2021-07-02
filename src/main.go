package main

import (
	"github.com/KaoruOhbayashi/golang_echo/dao"
	"github.com/KaoruOhbayashi/golang_echo/tao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

func main() {

	//Connect mysql
	db := dao.GormConnect()
	defer db.Close()

	e := echo.New()

	//Road teplates
	t := &tao.Template{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/index") // Redirect to home
	})

	//Output articles
	e.GET("index", func(c echo.Context) error{
		articles := dao.GetArticles(db)
		return  c.Render(http.StatusOK, "index", articles)
	})

	//Post article
	e.POST("/insert", func(c echo.Context) error{
		dao.InsertArticle(db,c.FormValue("title"),c.FormValue("contents"))
		return c.Redirect(http.StatusFound, "/index") // Redirect to home
	})

	e.Logger.Fatal(e.Start(":8080"))
}
