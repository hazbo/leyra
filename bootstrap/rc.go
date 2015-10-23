package bootstrap

import (
	"io/ioutil"
	"os"

	"gopkg.in/leyra/gorm.v1"
	_ "gopkg.in/leyra/mysql.v1"
	"gopkg.in/leyra/toml.v1"
)

// RcConfig is the base struct for how our configuration is all organised from
// the ./etc/rc.conf file. Anything in here is linked to the application
// runtime.
type RcConfig struct {
	Database struct {
		EnableDatabase string
		Database       string
		AutoMigrate    string
		Mysql          MysqlDatabase
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
	db, err := gorm.Open("mysql", "root@/leyra?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
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
