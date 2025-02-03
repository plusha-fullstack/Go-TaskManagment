package api

import (
	"net/http"
	"task-tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler contains the database instance
type UserHandler struct {
	DB *gorm.DB
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User
	h.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
