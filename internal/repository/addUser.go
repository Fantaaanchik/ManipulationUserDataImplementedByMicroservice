package repository

import (
	"log"
	"testproject/internal/db"
	"testproject/models"
)

func (r *Repository) AddNewUserToDB(users models.User) error {
	err := db.GetDB().Create(&users).Error
	if err != nil {
		db.GetDB().Rollback()
		log.Printf("err: %d", err)
		return err
	}

	return nil
}
