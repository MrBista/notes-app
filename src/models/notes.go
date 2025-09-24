package models

import (
	"time"
)

type Note struct {
	Id         int       `gorm:"primary_key;column:id;autoIncrement"`
	Status     int       `gorm:"column:status"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	IsFavorite bool      `gorm:"column:is_favorite"`
	ReminderAt time.Time `gorm:"column:reminder_at"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (n *Note) TableName() string {
	return "notes"
}
