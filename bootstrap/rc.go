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

package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"

	//	_ "gopkg.in/leyra/go-sqlite3.v1"
	"gopkg.in/leyra/gorm.v1"
	_ "gopkg.in/leyra/mysql.v1"
	"gopkg.in/leyra/toml.v1"
)

// RcConfig is the base struct for how our configuration is all organised from
// the ./etc/rc.conf file. Anything in here is linked to the application
// runtime.
type RcConfig struct {
	Application struct {
		Key string
	}

	Database struct {
		EnableDatabase string
		Database       string
		AutoMigrate    string
		Mysql          MysqlDatabase
		Sqlite         SqliteDatabase
	}

	Server struct {
		Port string
	}
}

// MysqlDatabase represents the basic configuration of how the user can connect
// to this particular kind of database.
type MysqlDatabase struct {
	Username string
	Password string
	Port     string
	Name     string
}

type SqliteDatabase struct {
	DbPath string
}

// NewRcConfig returns an empty instance of *RcConfig
func NewRcConfig() *RcConfig {
	return new(RcConfig)
}

// Apply takes the configuration from ./etc/rc.conf and applys each aspect of it
// to a newly create instance of *RcConfig.
func (c *RcConfig) Apply() *RcConfig {
	buf := configBuffer("./etc/rc.conf")

	if err := toml.Unmarshal(buf, c); err != nil {
		panic(err)
	}
	return c
}

// Connect is currently code for testing / debugging the use of gorm.DB
// This will not remain.
func (c *RcConfig) Connect() gorm.DB {
	// Just code for testing DB connections etc...
	var db gorm.DB
	var err error

	// Connect to a MySQL database
	if c.Database.Database == "mysql" {
		conn := fmt.Sprintf(
			"%s@/%s?charset=utf8&parseTime=True&loc=Local",
			c.Database.Mysql.Username,
			c.Database.Mysql.Name,
		)

		db, err = gorm.Open("mysql", conn)

		if err != nil {
			panic(err)
		}
	}

	// Create or use a SQLite database
	if c.Database.Database == "sqlite" {
		//		db, err = gorm.Open("sqlite3", c.Database.Sqlite.DbPath)
	}

	return db
}

func configBuffer(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return buf
}
