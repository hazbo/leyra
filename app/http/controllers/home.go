package controller

import (
	"gopkg.in/leyra/echo.v1"
	"net/http"
)

type Home struct {
}

func NewHomeController() Home {
	return Home{}
}

func (h Home) Home(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
