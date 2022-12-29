package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(g *echo.Group) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		// JWT の署名を行うために使用する秘密鍵を保持
		SigningKey: []byte("mySecret"),

		// JWT署名アルゴリズム
		SigningMethod: "HS256",
	}))
}
