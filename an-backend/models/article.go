package models

import "time"

type Article struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser    int       `json:"id_user" gorm:"not null"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Cover     string    `json:"cover" gorm:"type:varchar(255)"`
	Author    string    `json:"author" gorm:"type:varchar(100);not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Summary   string    `json:"summary" gorm:"type:varchar(500)"`
	Category  string    `json:"category" gorm:"type:varchar(50)"`
	Tags      string    `json:"tags" gorm:"type:varchar(255)"`
	Likes     int       `json:"likes" gorm:"default:0"`
	Favorites int       `json:"favorites" gorm:"default:0"`
	Views     int       `json:"views" gorm:"default:0"`
	Status    string    `json:"status" gorm:"type:enum('draft','published','archived');default:'draft'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ArticleRequest struct {
	IDUser   int    `json:"id_user" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Cover    string `json:"cover"`
	Author   string `json:"author" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Summary  string `json:"summary"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
}
