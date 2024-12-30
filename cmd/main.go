package main

import (
	"log"

	"github.com/jackcooperusesvim/coopGo/handler"
	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct{}

func main() {

	err := model.CreateTables()

	if err == nil {
		err = model.BuildTables()

		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}

	app := echo.New()

	//Various Middlewares
	app.Use(middleware.CSRF())
	app.Use(middleware.Secure())
	app.Use(middleware.Logger())

	courseHandler := handler.CourseHandler{}

	app.GET("/course", courseHandler.HandleCourseShow)
	app.GET("/course/edit/:id", courseHandler.HandleCourseEdit)
	app.GET("/course/new", courseHandler.HandleCourseNew)

	app.POST("/course/update", courseHandler.HandleCoursePost)
	app.POST("/course/create", courseHandler.HandleCourseCreate)
	app.POST("/course/delete", courseHandler.HandleCourseDelete)

	log.Println("app created")
	err = app.Start("localhost:4321")
	log.Println("app ended on error:")
	log.Println(err)
}
