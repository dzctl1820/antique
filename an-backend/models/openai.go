package models

import "time"

type Assistant struct {
	UserContent      string
	AssistantContent string
}

type TempAssistant struct {
	Role    string `gorm:"type:ENUM('system', 'user', 'assistant');not null;column:role" json:"role"`
	Content string `gorm:"type:text;not null;column:content" json:"content"`
}

type AiChat struct {
	Model            string      `gorm:"type:varchar(32);not null;column:model" json:"model"`
	SystemPromot     string      `gorm:"type:text;not null;column:system_prompt" json:"system"`
	UserContent      string      `gorm:"type:text;not null;column:user_content" json:"user"`
	AssistantContent []Assistant `gorm:"-" json:"assistant"`
	SessionId        int64       `gorm:"type:bigint;not null;column:session_id;index" json:"session_id"`
}

type ChatConversation struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	SessionID int64     `gorm:"type:bigint;not null;column:session_id;index" json:"session_id"`
	Role      string    `gorm:"type:ENUM('system', 'user', 'assistant');not null;column:role" json:"role"`
	Content   string    `gorm:"type:text;not null;column:content" json:"content"`
	Model     string    `gorm:"type:varchar(32);column:model" json:"model"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
}

type ChatSession struct {
	SessionID int        `gorm:"primaryKey;autoIncrement;column:session_id" json:"session_id"`
	UserID    int        `gorm:"type:int;column:user_id;index" json:"user_id"`
	Title     string     `gorm:"type:varchar(255);column:title" json:"title"`
	CreatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
	ExpiresAt *time.Time `gorm:"type:timestamp;column:expires_at" json:"expires_at"`
}
