package handler

import (
	"github.com/jackcooperusesvim/coopGo/view/auth"
	"github.com/labstack/echo/v4"
	"log"
)

type AuthHandler struct {
	PermissionGroups []string
}

func (h AuthHandler) AuthPage(c echo.Context) error {
	log.Println("AuthPage")
	return render(c, auth.LoginPage())
}

func (h AuthHandler) Auth(c echo.Context) error {
	log.Println(c.FormParams())
	log.Println(h.PermissionGroups)
	return nil
}
