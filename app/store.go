package app

import (
	"gopkg.in/leyra/gorm.v1"
)

// Store is a place where you can pub various dependencies for easy access later
// on. By default an instance of gorm.DB will be put in here providing that
// enable_database = "YES" in your ../etc/rc.conf file.
type Store struct {
	DB gorm.DB
}

// S provides a shortcut way of accessing the store from inside your
// application.
var S Store
