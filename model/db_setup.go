package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func CreateDatabase() {

	dsn := "host=localhost user=postgres password=password dbname=students port=5432"
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to database xxx")
	}

	// register Student model
	err = database.AutoMigrate(&Student{})
	if err != nil {
		return
	}
	Db = database
}
