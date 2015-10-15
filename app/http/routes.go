package http

import (
	"github.com/labstack/echo"

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
