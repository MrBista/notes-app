package models

import "time"

type User struct {
	Id        int       `gorm:"primary_key;column:id;autoIncrement"`
	FullName  string    `gorm:"column:full_name"`
	Password  string    `gorm:"column:password"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}
