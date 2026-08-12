package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(dbString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	client, _ := db.DB()

	client.SetMaxIdleConns(5)
	client.SetMaxOpenConns(10)
	client.SetConnMaxLifetime(time.Hour)

	return db
}
