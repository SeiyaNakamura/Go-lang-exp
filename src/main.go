package main

import (
	"github.com/KaoruOhbayashi/golang_echo/dao"
	"github.com/KaoruOhbayashi/golang_echo/tao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func main() {

	db := dao.GormConnect()
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	t := &tao.Template{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/index", tao.Hello)

	e.Logger.Fatal(e.Start(":8080"))
}
