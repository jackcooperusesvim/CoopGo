package middleware

import (
	"errors"

	"github.com/jackcooperusesvim/coopGo/model"
	"github.com/labstack/echo/v4"
	"log"
)

type ACL struct {
	AuthGroups []string `json:"statuses"`
}

func NewACL() *ACL {
	return &ACL{
		AuthGroups: []string{"all"},
	}
}

func (m *ACL) Restrict(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		priv := c.Get("privledge_level").(string)
		log.Println(priv)
		for _, group := range m.AuthGroups {
			if group == priv {
				return next(c)
			}
		}
		return errors.New("This user type is not allowed to access this endpoint")
	}
}

func BehindAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token_cookie, err := c.Cookie("session_token")
		log.Println(token_cookie)
		token := token_cookie.Value

		priv, account_id, err := model.ValidateToken(token)
		log.Println(priv)

		if err != nil {
			log.Println(err)
			return c.NoContent(401)
		}

		c.Set("privledge_level", priv)
		c.Set("account_id", account_id)

		return next(c)

	}
}
