package main

import (
	"log"

	"github.com/jackcooperusesvim/coopGo/handler"
	cm "github.com/jackcooperusesvim/coopGo/middleware"
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
	app.Use(middleware.Logger())
	app.Use(middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			Skipper:     middleware.DefaultSkipper,
			TokenLength: 32,
			TokenLookup: "form:csrf",
			ContextKey:  "csrf",
		}))
	app.Use(middleware.Secure())

	courseHandler := handler.CourseHandler{}

	AuthHandler := handler.AuthHandler{}
	// adminACL := &cm.ACL{
	// 	AuthGroups: []string{"admin"},
	// }

	app.GET("/login", AuthHandler.AuthPage)
	app.POST("/family/login", AuthHandler.Login)
	app.POST("/admin/login", AuthHandler.Login)

	app.GET("/course", cm.BehindAuth(courseHandler.HandleCourseShow))
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
