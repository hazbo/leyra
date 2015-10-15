package http

import (
	"gopkg.in/leyra/echo.v1"

	"leyra/app/http/controllers"
)

func InitRoutes() {
	e := echo.New()
	e.Get("/", routeHome)

	e.Run(":3000")
}

func routeHome(c *echo.Context) error {
	return controller.NewHomeController().Home(c)
}
