package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
