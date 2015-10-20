package bootstrap

import (
	"gopkg.in/leyra/gorm.v1"
	_ "gopkg.in/leyra/mysql.v1"
	//"leyra/app/models"
	"gopkg.in/leyra/toml.v1"
	bootutil "leyra/bootstrap/util"
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

func (dc *DatabaseConfig) Apply() *DatabaseConfig {
	buf := bootutil.ConfigBuffer("./etc/database.conf")

	if err := toml.Unmarshal(buf, dc); err != nil {
		panic(err)
	}

	return dc
}

func (dc *DatabaseConfig) Connect() gorm.DB {
	// Just code for testing DB connections etc...
	db, err := gorm.Open("mysql", "root@/leyra?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
	// db.AutoMigrate(&model.User{})
}
