package repository

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(cfg *viper.Viper) (db *gorm.DB) {
	var dbUser, dbPass, dbHost, dbPort, dbName string
	dbUser = cfg.GetString("database.username")
	dbPass = cfg.GetString("database.password")
	dbHost = cfg.GetString("database.host")
	dbPort = cfg.GetString("database.port")
	dbName = cfg.GetString("database.dbname")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	return db

}
