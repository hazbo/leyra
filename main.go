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

	e := http.Route()

	// Only attempt to make a database connection if it has been enabled in
	// etc/rc.conf
	if rc.Database.EnableDatabase == "YES" {
		// Load database settings from ./etc/database.conf
		db := rc.Connect()
		db.DB().Ping()
	}

	http.Serve(e, rc.Server.Port)
}
