package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	Name       string
	Note       string
	IsComplete bool
	CreatedAt  time.Time
}

// Todoテーブルのマイグレーション
func MigrateTodo(db *gorm.DB) error {
	err := db.AutoMigrate(&Todo{})
	return err
}
