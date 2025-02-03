package routes

import (
	"task-tracker/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userHandler := &api.UserHandler{DB: db}
	taskHandler := &api.TaskHandler{DB: db}

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTask)

	return r
}

// // Setup Routes THIS IS MAIN BELOW
// r := routes.SetupRouter(database)

// // Start Server
// port := "8080"
// fmt.Println("Server is running on port " + port)
// r.Run(":" + port)
