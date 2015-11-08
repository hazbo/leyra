package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
)

type Home struct {
}

func (h Home) Home(c *echo.Context) error {
	buff := new(bytes.Buffer)
	app.S.View.ExecuteTemplate(buff, "index.html", nil)

	return c.HTML(http.StatusOK, buff.String())
}
