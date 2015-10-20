package http

import (
	"gopkg.in/leyra/echo.v1"
)

func Serve(e *echo.Echo, port string) {
	e.Run(port)
}
