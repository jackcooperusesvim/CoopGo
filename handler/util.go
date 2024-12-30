package handler

import (
	"time"

	"github.com/a-h/templ"
	"github.com/jackcooperusesvim/coopGo/view/layout"
	"github.com/labstack/echo/v4"
)

func valid_date(inp string) bool {
	_, err := time.Parse("yyyy-mm-dd", inp)
	if err != nil {
		return false
	}
	return true
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
func redirect(c echo.Context, url string) error {
	return render(c, layout.Redirect(url))
}
func redirect_home(c echo.Context) error {
	return render(c, layout.Redirect("/"))
}
