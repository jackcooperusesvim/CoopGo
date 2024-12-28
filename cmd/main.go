package main

import (
	"github.com/jackcooperusesvim/coopGo/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	courseHandler := handler.CourseHandler{}

	app.GET("/course", courseHandler.HandleCourseShow)

	app.Start("localhost:3000")
}
