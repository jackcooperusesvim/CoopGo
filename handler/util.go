package handler

import (
	"github.com/a-h/templ"
	"github.com/jackcooperusesvim/coopGo/view/layout"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
func redirect(c echo.Context, url string) error {
	return render(c, layout.Redirect(url))
}
func redirect_home(c echo.Context) error {
	return render(c, layout.Redirect("/"))
}
