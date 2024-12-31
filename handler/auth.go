package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/jackcooperusesvim/coopGo/view/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h *AuthHandler) AuthPage(c echo.Context) error {
	log.Println("AuthPage")
	return render(c, auth.LoginPage(c.Get("csrf").(string)))
}

func (h *AuthHandler) Login(c echo.Context) error {
	log.Println("AuthPage")
	email := c.FormValue("email")
	password := c.FormValue("password")
	token, _, _, err := model.Login(email, password)
	if err != nil {
		c.Response().Header().Set("HX-Redirect", "/course")
		return c.NoContent(401)
	}

	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().AddDate(0, 0, 1),
	})

	c.Response().Header().Set("HX-Redirect", "/course")
	//TODO: Send the user to the proper domain

	return c.NoContent(200)
}
