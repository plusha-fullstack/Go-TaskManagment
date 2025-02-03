package models

import (
	"gorm.io/gorm"
)

const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"

	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)

type Task struct {
	gorm.Model  // there ID and Create/deleted/updated at are wrapped in this struct
	Title       string
	Description string
	Status      string
	Priority    string
	AuthorID    uint // Foreign key
	Author      User `gorm:"foreignKey:AuthorID"`
	AssigneeID  uint // Foreign key
	Assignee    User `gorm:"foreignKey:AssigneeID"`
}
