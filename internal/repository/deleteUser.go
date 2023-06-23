package repository

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testproject/internal/db"
	"testproject/models"
)

func (user *UserRepository) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if result := db.GetDB().Delete(&models.User{}, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя из таблицы users удалены!"})
	return
}
