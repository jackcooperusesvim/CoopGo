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
		//TODO: ADD DEFAULT ADMIN

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
	pubAuthHandler := handler.AuthHandler{
		PermissionGroups: []string{"all"},
	}
	familyAuthHandler := handler.AuthHandler{
		PermissionGroups: []string{"family"},
	}
	adminAuthHandler := handler.AuthHandler{
		PermissionGroups: []string{"admin"},
	}

	app.GET("/login", pubAuthHandler.AuthPage)
	app.POST("/family/auth", familyAuthHandler.Auth)
	app.POST("/admin/auth", adminAuthHandler.Auth)

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
