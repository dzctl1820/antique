package service

import (
	"an-backend/models"
	"an-backend/utils"
	"gorm.io/gorm"
	"time"
)

type AiChatDB struct {
	DB *gorm.DB
}

func Chat(db *gorm.DB, userContent string, modelName string, systemPrompt string, assistantAndUser []models.Assistant, sessionId int64) string {

	content := utils.Chat(userContent, modelName, systemPrompt, assistantAndUser)
	if content == "" {
		return "无响应"
	}
	db.Create(&models.ChatConversation{
		SessionID: sessionId,
		Role:      "user",
		Content:   userContent,
		Model:     modelName,
	})
	db.Create(&models.ChatConversation{
		SessionID: sessionId,
		Role:      "assistant",
		Content:   content,
		Model:     modelName,
	})
	return content
}

func NewAiChatDB(db *gorm.DB, id int, title string) *models.ChatSession {
	aiChatSession := models.ChatSession{
		UserID:    id,
		Title:     "新的对话",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: &time.Time{},
	}
	if title != "" {
		aiChatSession.Title = title
	}
	db.Create(&aiChatSession)
	if db.Error != nil {
		return nil
	}
	return &aiChatSession
}

func GetAiChatDB(db *gorm.DB, id int64) []models.Assistant {
	var temp []models.TempAssistant
	if err := db.Table("chat_conversations").Where("session_id = ?", id).Order("created_at asc").Find(&temp).Error; err != nil {
		return nil
	}
	res := make([]models.Assistant, 0)

	userMessages := make([]string, 0)
	assistantMessages := make([]string, 0)

	for _, item := range temp {
		if item.Role == "user" {
			userMessages = append(userMessages, item.Content)
		} else if item.Role == "assistant" {
			assistantMessages = append(assistantMessages, item.Content)
		}
	}

	minLen := len(userMessages)
	if len(assistantMessages) < minLen {
		minLen = len(assistantMessages)
	}

	for i := 0; i < minLen; i++ {
		res = append(res, models.Assistant{
			UserContent:      userMessages[i],
			AssistantContent: assistantMessages[i],
		})
	}
	return res
}

func GetAiChatSessions(db *gorm.DB, userID int) []models.ChatSession {
	var sessions []models.ChatSession
	if err := db.Where("user_id = ?", userID).Order("created_at desc").Find(&sessions).Error; err != nil {
		return nil
	}
	return sessions
}

func DeleteAiChatDB(db *gorm.DB, sessionID string) {
	db.Delete(&models.ChatSession{}, "session_id = ?", sessionID)
	db.Delete(&models.ChatConversation{}, "session_id = ?", sessionID)
}

func UpdateAiChatDB(db *gorm.DB, sessionID string, title string) {
	db.Model(&models.ChatSession{}).Where("session_id = ?", sessionID).Update("title", title)
}

func GetAiChatConversationDB(db *gorm.DB, sessionID string) []models.ChatConversation {
	var conversations []models.ChatConversation
	if err := db.Where("session_id = ?", sessionID).Order("created_at asc").Find(&conversations).Error; err != nil {
		return nil
	}
	return conversations
}
