package main

import (
	"leyra/app/http"
	"leyra/bootstrap"
)

func main() {
	// Make calls to bootstrap here
	bootstrap.SetEnv()

	// Load database settings from etc/database.conf
	bootstrap.NewDatabaseConfig().Apply()

	http.InitRoutes()
}
