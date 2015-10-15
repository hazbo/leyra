package main

import (
	"leyra/app/http"
	"leyra/bootstrap"
)

func main() {
	// Make calls to bootstrap here
	bootstrap.SetEnv()

	http.InitRoutes()
}
