package main

import (
	"log"

	"github.com/jackcooperusesvim/coopGo/handler"
	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct{}

func main() {
	err := model.CreateTables()

	if err != nil {
		log.Println(err)
	}
	err = model.BuildTables()

	if err != nil {
		log.Println(err)
	}

	app := echo.New()

	courseHandler := handler.CourseHandler{}

	app.GET("/course", courseHandler.HandleCourseShow)

	app.Start("localhost:3000")
}
