package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/jackcooperusesvim/coopGo/model/sqlgen"
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
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("No id provided")
	}

	if err != nil {
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

func (h CourseHandler) HandleCourseDelete(c echo.Context) (err error) {
	log.Println("FormParams:")
	log.Println(c.FormParams())

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}

	q, ctx, err := model.DbInfo()

	if err != nil {
		return err
	}

	err = q.DeleteCourse(ctx, int64(id))

	c.Response().Header().Set("HX-Redirect", "/course")
	if err != nil {
		c.NoContent(http.StatusInternalServerError) // No body needed
		return err
	} else {
		return c.NoContent(http.StatusOK) // No body needed
	}

}

func (h CourseHandler) HandleCoursePost(c echo.Context) (err error) {

	log.Println("FormParams:")
	log.Println(c.FormParams())

	q, ctx, err := model.DbInfo()

	if err != nil {
		return err
	}

	ucp := sqlgen.UpdateCourseParams{}

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}
	ucp.ID = int64(id)
	ucp.Name = c.FormValue("name")
	ucp.Desc = c.FormValue("desc")
	ucp.StartDate = c.FormValue("start_date")
	ucp.EndDate = c.FormValue("end_date")

	crs, err := q.UpdateCourse(ctx, ucp)
	log.Println(crs.EndDate)

	c.Response().Header().Set("HX-Redirect", "/course")
	if err != nil {
		c.NoContent(http.StatusInternalServerError) // No body needed
		return err
	} else {
		return c.NoContent(http.StatusOK) // No body needed
	}
}

func (h CourseHandler) HandleCourseNew(c echo.Context) error {
	crs := sqlgen.Course{
		ID:        -123,
		Name:      "",
		Desc:      "",
		StartDate: "",
		EndDate:   "",
	}
	return render(c, course.New(crs))
}
func (h CourseHandler) HandleCourseCreate(c echo.Context) error {
	ucp := sqlgen.CreateCourseParams{}
	q, ctx, err := model.DbInfo()

	if err != nil {
		return err
	}

	ucp.Name = c.FormValue("name")
	ucp.Desc = c.FormValue("desc")
	ucp.StartDate = c.FormValue("start_date")
	ucp.EndDate = c.FormValue("end_date")

	_, err = q.CreateCourse(ctx, ucp)

	c.Response().Header().Set("HX-Redirect", "/course")

	if err != nil {
		c.NoContent(http.StatusInternalServerError) // No body needed
		return err
	} else {
		return c.NoContent(http.StatusOK) // No body needed
	}
}
