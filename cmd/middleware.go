package main

import (
	"errors"
	"sync"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/labstack/echo/v4"
)

type ACL struct {
	AllowedGroups []string `json:"statuses"`
	mutex         sync.RWMutex
}

func AllowAccess(handler func([]string, echo.Context) error) {

}
func NewACL() *ACL {
	return &ACL{
		AllowedGroups: []string{"all"},
	}
}

func (m *ACL) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for i := range m.AllowedGroups {
			if i == c.Get("privledge_level") {
				return next(c)
			}
		}
		return errors.New("This user type is not allowed to access this endpoint")
	}
}

type Auth struct {
	mutex sync.RWMutex
}

func NewAuth() *Auth {
	return &Auth{}
}

func (m *Auth) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("session_token").(string)
		priv, account_id, err := model.ValidateToken(token)
		if err != nil {
			return err
		}

		c.Set("privledge_level", priv)
		c.Set("account_id", account_id)

		return next(c)

	}
}
