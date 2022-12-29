package middlewares

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {
			//　strings.Contains は、文字列が特定の文字列を含んでいるかどうかを判断する
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "not cookie")
			}
		}

		if cookie.Value == "jon" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}
