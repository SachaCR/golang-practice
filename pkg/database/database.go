package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() *gorm.DB {
	var host = viper.GetString("postgres.host")
	var port = viper.GetString("postgres.port")
	var user = viper.GetString("postgres.user")
	var password = viper.GetString("postgres.password")
	var dbName = viper.GetString("postgres.database")

	var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Paris", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
