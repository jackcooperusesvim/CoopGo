package handler

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/jackcooperusesvim/coopGo/view/course"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct{}

func (h CourseHandler) HandleCourseShow(c echo.Context) error {
	q, ctx, err := model.DbInfo()

	if err != nil {
		return err
	}

	courses, err := q.ListCourse(ctx)

	if err != nil {
		return err
	}
	return render(c, course.List(courses))
}

func (h CourseHandler) HandleCourseEdit(c echo.Context) error {
	log.Println("HandleCourseEdit")
	var id int = -1
	var err error = nil

	names := c.Param("id")
	if names != "" {
		split_id := strings.Split(c.Param("id"), "=")

		if len(split_id) == 2 {
			id, err = strconv.Atoi(split_id[1])
			log.Println("id:")
			log.Println(id)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("No id provided")
	}
	log.Println("final_id:")
	log.Println(id)

	if err != nil {
		log.Println(err)
		return err
	}

	q, ctx, err := model.DbInfo()

	if err != nil {
		return err
	}

	rel_course, err := q.GetCourse(ctx, int64(id))

	if err != nil {
		return err
	}
	return render(c, course.Edit(rel_course))
}

type CourseForm struct {
	id        string `form:"id"`
	Name      string `form:"name"`
	Desc      string `form:"desc"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

func (h CourseHandler) HandleCoursePost(c echo.Context) (err error) {

	r := c.Request()
	log.Println(r)

	err = r.ParseForm()
	if err != nil {
		return err
	}
	log.Println(r.ParseForm())

	cf := new(CourseForm)
	if err := c.Bind(cf); err != nil {
		return err
	}
	log.Println(cf)

	c.Response().Header().Set("HX-Redirect", "/course")

	return redirect_home(c)
}
