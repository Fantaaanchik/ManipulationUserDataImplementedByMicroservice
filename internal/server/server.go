package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testproject/internal/service"
	"testproject/models"
)

type Handler struct {
	Engine   *gin.Engine
	services *service.Services
}

func NewH(services *service.Services, engine *gin.Engine) *Handler {
	return &Handler{services: services, Engine: engine}
}

func (h Handler) AllRoutes() {

	//h.Engine.POST("/add_users", h.AddNewUser)
	//r.PUT("/change_users/:id", ur.ChangeUsers)
	//r.DELETE("/delete_users/:id", ur.DeleteUsers)

	h.Engine.GET("/get_users", h.GetUsers)
	h.Engine.POST("/add_users", h.AddUsers)
	h.Engine.PUT("/update_user_data/:id", h.UpdateUserData)
	h.Engine.DELETE("/delete_user_data/:id", h.DeleteUserData)

}

func (h Handler) GetUsers(c *gin.Context) {

	user, err := h.services.GetUserFromDB()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": user})
}

func (h Handler) AddUsers(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.services.AddNewUserToDB(user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Новый пользователь добавлен"})
}

func (h Handler) UpdateUserData(c *gin.Context) {

	userID := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := h.services.UpdateUserDataFromDB(userID, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении данных пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя обновлены"})
}

func (h Handler) DeleteUserData(c *gin.Context) {

	userID := c.Param("id")

	// Удаление пользователя по ID
	if err := h.services.DeleteUserDataFromDB(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя удалены"})
}
