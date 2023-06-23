package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testproject/models"
)

var db *gorm.DB

func ConnectionToDB() *gorm.DB {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to db, err: ", err.Error())
		return nil
	}

	err = db.AutoMigrate(&models.User{})

	return db
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	PgDB, err := db.DB()
	err = PgDB.Close()
	if err != nil {
		log.Println("не удалось закрыть DB: ", err.Error())
	}
}
