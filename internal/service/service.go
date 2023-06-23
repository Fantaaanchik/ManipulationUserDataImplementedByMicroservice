package service

import (
	"fmt"
	"testproject/internal/repository"
	"testproject/models"
)

type Services struct {
	Repository *repository.Repository
}

func NewService(repo *repository.Repository) *Services {
	return &Services{Repository: repo}
}

func (s Services) GetUserFromDB() ([]models.User, error) {
	return s.Repository.GetUserFromDB()
}

func (s Services) AddNewUserToDB(user models.User) error {
	return s.Repository.AddNewUserToDB(user)
}

func (s Services) UpdateUserDataFromDB(userID string, user models.User) error {
	existingUser, err := s.Repository.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("не удалось получить пользователя с идентификатором %s: %w", userID, err)
	}

	existingUser.Fio = user.Fio
	existingUser.Number = user.Number

	err = s.Repository.UpdateUserDataFromDB(existingUser)
	if err != nil {
		return fmt.Errorf("не удалось сохранить обновления пользователя: %w", err)
	}

	return nil
}

func (s Services) DeleteUserDataFromDB(userID string) error {
	err := s.Repository.DeleteUserByID(userID)
	if err != nil {
		return fmt.Errorf("не удалось получить пользователя с идентификатором %s: %w", userID, err)
	}

	return nil
}
