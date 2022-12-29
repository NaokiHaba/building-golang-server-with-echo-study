package api

import (
	"github.com/labstack/echo"
	"github.com/naokis-practice-project/practice_echo_example/src/api/handlers"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
