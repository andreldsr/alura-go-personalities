package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=personalities port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}

}
