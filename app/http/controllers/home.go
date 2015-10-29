package controller

import (
	"net/http"

	"gopkg.in/leyra/echo.v1"
)

type Home struct {
}

func (h Home) Home(c *echo.Context) error {
	return c.HTML(http.StatusOK, "Hello, World!")
}
