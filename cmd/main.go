package main

import (
	"log"
	"time"

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

	//Background processes
	go func() {
		q, ctx, err := model.DbInfo()
		if err == nil {
			for {
				err := q.PubliclyUnaliveTokens(ctx)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(time.Minute)
			}
		}
		log.Println(err)
	}()

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
	adminACL := &cm.ACL{
		AuthGroups: []string{"admin"},
	}

	app.GET("/login", AuthHandler.AuthPage)
	app.POST("/new_session", AuthHandler.Login)

	app.GET("/course", cm.BehindAuth(courseHandler.HandleCoursePage))
	app.GET("/course/edit/:id", cm.BehindAuth(adminACL.Restrict(courseHandler.HandleCourseEdit)))
	app.GET("/course/new", cm.BehindAuth(adminACL.Restrict(courseHandler.HandleCourseNew)))

	app.POST("/course/update", cm.BehindAuth(adminACL.Restrict(courseHandler.HandleCoursePost)))
	app.POST("/course/create", cm.BehindAuth(adminACL.Restrict(courseHandler.HandleCourseCreate)))
	app.POST("/course/delete", cm.BehindAuth(adminACL.Restrict(courseHandler.HandleCourseDelete)))

	log.Println("app created")
	err = app.Start("localhost:4321")
	log.Println("app ended on error:")
	log.Println(err)
}

func GETBehindAuth(app *echo.Echo, route string, handler echo.HandlerFunc, allowed_groups []string) {
	acl := cm.ACL{
		AuthGroups: allowed_groups}

	app.GET(route, cm.BehindAuth(acl.Restrict(handler)))
}
func POSTBehindAuth(app *echo.Echo, route string, handler echo.HandlerFunc, allowed_groups []string) {
	acl := cm.ACL{
		AuthGroups: allowed_groups}

	app.POST(route, cm.BehindAuth(acl.Restrict(handler)))
}
