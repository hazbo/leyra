package http

import (
	"gopkg.in/leyra/echo.v1"
)

// Serve takes echo and a port formatted as a string like so ":3000" to then
// start the server. This port is passed through from main.go which in turn is
// taken from rc.conf
func Serve(e *echo.Echo, port string) {
	e.Run(port)
}
