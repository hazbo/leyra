package bootstrap

import (
	//"gopkg.in/leyra/gorm.v1"
	//_ "gopkg.in/leyra/mysql.v1"
	//"leyra/app/models"
	"gopkg.in/leyra/toml.v1"
	"io/ioutil"
	"os"
)

type DatabaseConfig struct {
	Database    string
	AutoMigrate string

	Mysql struct {
		Username string
		Password string
		Port     string
	}
}

func NewDatabaseConfig() *DatabaseConfig {
	return new(DatabaseConfig)
}

func (dc *DatabaseConfig) Apply() {
	f, err := os.Open("./etc/database.conf")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, dc); err != nil {
		panic(err)
	}
}

func DatabaseConnect() {
	// Just code for testing DB connections etc...
	// db, err := gorm.Open("mysql", "root@/leyra?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// db.DB()
	// db.DB().Ping()
	// db.AutoMigrate(&model.User{})
}
