package models

import "time"

type Todo struct {
	ID         uint `gorm:"primaryKey;default:auto_random()"`
	Name       string
	Note       string
	IsComplete bool
	CreatedAt  time.Time
}
