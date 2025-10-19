package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Avatar    string
	Phone     string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
