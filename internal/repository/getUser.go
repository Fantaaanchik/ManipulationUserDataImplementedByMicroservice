package repository

import (
	"gorm.io/gorm"
	"log"
	"testproject/internal/db"
	"testproject/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	UserRep models.User
}

func (r *Repository) GetUserFromDB() ([]models.User, error) {
	var users []models.User

	err := db.GetDB().Find(&users).Error
	if err != nil {
		db.GetDB().Rollback()
		log.Printf("err: %d", err)
		return nil, err
	}

	return users, nil
}
