package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "junaid:1234@/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// Create the database if it doesn't exist
	err = d.Exec("CREATE DATABASE IF NOT EXISTS gobookstore").Error
	if err != nil {
		panic(err)
	}

	// Close the initial connection
	d.Close()

	// Reconnect to the MySQL server with the newly created database
	d, err = gorm.Open("mysql", "junaid:1234@/gobookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
