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

// The main package is responsible for bootstrapping the application and dealing
// with any runtime configuration. There are both calles to app.Before and
// app.After() that will allow you to run any code either just before or after
// the web server has started.
//
// Typically you shouldn't have to edit anything in this file for now.

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"leyra/app"
	"leyra/app/http"
	"leyra/bootstrap"
)

func main() {
	// Create application WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

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

		app.S.DB = db
	}

	// Parse and cache all the templates here, ready to go into the store
	templates := template.New("template")

	filepath.Walk("./app/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			templates.ParseFiles(path)
		}

		return nil
	})

	app.S.Template = templates

	// Start application web server
	app.Before()
	go http.Serve(e, rc.Server.Port)
	app.After()

	wg.Wait()
}
