package app

import (
	"html/template"

	"gopkg.in/leyra/echo.v1"
	"gopkg.in/leyra/gorm.v1"
	"gopkg.in/leyra/sessions.v1"
)

// Store is a place where you can pub various dependencies for easy access later
// on. By default an instance of gorm.DB will be put in here providing that
// enable_database = "YES" in your ../etc/rc.conf file.
type Store struct {
	DB      gorm.DB
	View    *template.Template
	Session *sessions.FilesystemStore

	Set func(c *echo.Context, key string, value interface{})
	Get func(c *echo.Context, key string) interface{}
}

// S provides a shortcut way of accessing the store from inside your
// application.
var S Store
