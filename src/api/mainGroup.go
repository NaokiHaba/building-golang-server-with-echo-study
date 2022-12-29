package api

import (
	"github.com/labstack/echo"
	"github.com/naokis-practice-project/practice_echo_example/src/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Hello)
	e.GET("/login", handlers.Login)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)

	e.POST("/dogs", handlers.AddDog)

	e.POST("/hamsters", handlers.AddHamster)
}
