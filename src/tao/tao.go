package tao

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

