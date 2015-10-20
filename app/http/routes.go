package http

import (
	"gopkg.in/leyra/echo.v1"
	"leyra/app/http/controllers"
)

func Route() *echo.Echo {
	e := echo.New()
	e.Get("/", routeHome)

	return e
}

func routeHome(c *echo.Context) error {
	return controller.NewHomeController().Home(c)
}
