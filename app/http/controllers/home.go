package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type Home struct {
}

func NewHomeController() Home {
	return Home{}
}

func (h Home) Home(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
