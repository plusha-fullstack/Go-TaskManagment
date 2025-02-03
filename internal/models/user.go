package models

import _ "gorm.io/gorm"

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Email    *string
	Role     string `gorm:"default:'user';not null" json:"role"`
}
