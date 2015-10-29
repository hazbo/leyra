package http

import (
	"gopkg.in/leyra/echo.v1"

	"leyra/app/http/controllers"
)

// Route currently creates a new instance of echo and attaches routes to
// patterns that can be defined in this file. I'm still not so sure about how
// all of this should work - for now it's fine though.
func Route() *echo.Echo {
	e := echo.New()
	e.Get("/", routeHome)

	return e
}

func routeHome(c *echo.Context) error {
	return new(controller.Home).Home(c)
}
