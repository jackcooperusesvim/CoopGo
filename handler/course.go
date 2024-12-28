package handler

import (
	"log"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/jackcooperusesvim/coopGo/view/course"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct{}

func (h CourseHandler) HandleCourseShow(c echo.Context) error {
	q, ctx, err := model.DbInfo()

	if err != nil {
		log.Println(err)
		return err
	}

	courses, err := q.ListCourse(ctx)

	if err != nil {
		log.Println(err)
		return err
	}
	return render(c, course.List(courses))
}
