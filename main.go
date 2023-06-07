package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", helloword)
	e.GET("/", home)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloword(c echo.Context) error {
	return c.String(http.StatusOK, "helloworld")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		// fmt.Println("tidak ada")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
