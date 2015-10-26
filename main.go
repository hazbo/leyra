// leyra v0.0.1-dev
//
// (c) Ground Six 2015
//
// @package leyra
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"leyra/app"
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

	app.Start()
	http.Serve(e, rc.Server.Port)
}
