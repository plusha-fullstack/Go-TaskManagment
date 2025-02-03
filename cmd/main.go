package main

import (
	"task-tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=pass dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	email := "irina_vtn@yandex.ru"
	// Create
	db.Create(&models.User{Username: "Irina", Email: &email, Password: "123456", Role: "admin"})

	// Read
	var User models.User
	db.First(&User, 1) // find product with integer primary key

	db.Model(&User).Update("Username", "George")
	// Update - update multiple fields
	db.Model(&User).Updates(&models.User{Username: "Wakik", Email: &email, Password: "654321", Role: "user"}) // non-zero fields
	db.Model(&User).Updates(map[string]interface{}{"Username": "Wakik", "Email": &email, "Password": "6ya gay", "Role": "guest"})

	// Delete - delete product
	//db.Delete(&User, 1)
	// all this finally work!!
}
