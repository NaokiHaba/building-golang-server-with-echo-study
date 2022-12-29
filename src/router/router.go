package router

import (
	"github.com/labstack/echo"
	"github.com/naokis-practice-project/practice_echo_example/src/api"
	"github.com/naokis-practice-project/practice_echo_example/src/api/middlewares"
)

func New() *echo.Echo {
	// create echo instance
	e := echo.New()

	// create route groups
	ag := e.Group("/admin")
	cg := e.Group("/cookie")
	jg := e.Group("/jwt")

	// set middlewares
	middlewares.SetMainMiddleware(e)
	middlewares.SetAdminMiddlewares(ag)
	middlewares.SetCookieMiddlewares(cg)
	middlewares.SetJwtMiddlewares(jg)

	// set routers
	api.MainGroup(e)
	api.AdminGroup(ag)
	api.CookieGroup(cg)
	api.JwtGroup(jg)

	return e
}
