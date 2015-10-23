package main

import (
	"leyra/app/http"
	"leyra/bootstrap"
)

func main() {
	// Make calls to bootstrap here
	bootstrap.SetEnv()

	// Runtime configuration
	rc := bootstrap.NewRcConfig()
	rc.Apply()

	// Server configuration
	sc := bootstrap.NewServerConfig()
	sc.Apply()

	e := http.Route()

	// Only attempt to make a database connection if it has been enabled in
	// etc/rc.conf
	if rc.DatabaseEnable == "YES" {
		// Load database settings from ./etc/database.conf
		dc := bootstrap.NewDatabaseConfig().Apply()
		db := dc.Connect()

		db.DB().Ping()
	}

	http.Serve(e, sc.Port)
}
