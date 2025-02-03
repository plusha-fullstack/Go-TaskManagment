package api

import (
	"net/http"
	"task-tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TaskHandler contains the database instance
type TaskHandler struct {
	DB *gorm.DB
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// GetTasks handles GET /tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []models.Task
	h.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// GetTask handles GET /tasks/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
