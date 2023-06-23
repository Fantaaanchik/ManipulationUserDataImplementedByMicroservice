package repository

import (
	"log"
	"testproject/internal/db"
	"testproject/models"
)

func (r *Repository) UpdateUserDataFromDB(users *models.User) error {
	err := db.GetDB().Updates(&users).Error
	if err != nil {
		db.GetDB().Rollback()
		log.Printf("err: %d", err)
		return err
	}

	return nil
}

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) DeleteUserByID(id string) error {
	err := db.GetDB().Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
