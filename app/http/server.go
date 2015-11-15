package http

import (
	"gopkg.in/leyra/echo.v1"
	"gopkg.in/leyra/grace.v1/gracehttp"
)

// Serve takes echo and a port formatted as a string like so ":3000" to then
// start the server. This port is passed through from main.go which in turn is
// taken from rc.conf
func Serve(e *echo.Echo, port string) {
	s := e.Server(port)
	s.TLSConfig = nil

	e.Static("/static", "public")

	gracehttp.Serve(s)
}
