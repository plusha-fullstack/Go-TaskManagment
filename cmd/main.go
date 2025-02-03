package main

import (
	"fmt"
	"log"
	"task-tracker/db"
	"task-tracker/routes"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	r := routes.SetupRouter(database)

	// Start Server
	port := "8080"
	fmt.Println("Server is running on port " + port)
	r.Run(":" + port)
	// // Migrate the schema

	// email := "irina_vtn@yandex.ru"
	// // Create
	// db.Create(&models.User{Username: "Irina", Email: &email, Password: "123456", Role: "admin"})

	// // Read
	// var User models.User
	// db.First(&User, 1) // find product with integer primary key

	// db.Model(&User).Update("Username", "George")
	// // Update - update multiple fields
	// db.Model(&User).Updates(&models.User{Username: "Wakik", Email: &email, Password: "654321", Role: "user"}) // non-zero fields
	// db.Model(&User).Updates(map[string]interface{}{"Username": "Wakik", "Email": &email, "Password": "6ya gay", "Role": "guest"})

	// Delete - delete product
	//db.Delete(&User, 1)
	// all this finally work!!
}
