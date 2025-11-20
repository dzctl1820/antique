package handler

import (
	"an-backend/internal/service"
	"an-backend/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AiChatService struct {
	DB *service.AiChatDB
}

func (s *AiChatService) Chat(c *gin.Context) {
	var chat models.AiChat
	chat.Model = "qwen-plus"
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	chat.AssistantContent = service.GetAiChatDB(s.DB.DB, chat.SessionId)
	content := service.Chat(s.DB.DB, chat.UserContent, chat.Model, chat.SystemPromot, chat.AssistantContent, chat.SessionId)
	c.JSON(200, gin.H{"content": content})
}

func (s *AiChatService) NewSession(c *gin.Context) {

	var sessionReg models.ChatSession
	err := c.ShouldBindJSON(&sessionReg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, ok := c.Get("userID")
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(400, gin.H{"error": "userID is not uint"})
		return
	}
	sessionReg.UserID = int(uid)
	session := service.NewAiChatDB(s.DB.DB, sessionReg.UserID, sessionReg.Title)
	c.JSON(200, gin.H{"session_id": session.SessionID, "content": "创建成功"})
}

func (s *AiChatService) GetSession(c *gin.Context) {
	userID, ok := c.Get("userID")
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(400, gin.H{"error": "userID is not uint"})
		return
	}
	sessions := service.GetAiChatSessions(s.DB.DB, int(uid))
	c.JSON(200, sessions)
}

func (s *AiChatService) DeleteSession(c *gin.Context) {
	var sessionReg models.ChatSession
	err := c.ShouldBindJSON(&sessionReg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	service.DeleteAiChatDB(s.DB.DB, strconv.Itoa(sessionReg.SessionID))
	c.JSON(200, gin.H{"content": "删除成功"})
}

func (s *AiChatService) UpdateSession(c *gin.Context) {
	var sessionReg models.ChatSession
	err := c.ShouldBindJSON(&sessionReg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	service.UpdateAiChatDB(s.DB.DB, strconv.Itoa(sessionReg.SessionID), sessionReg.Title)
	c.JSON(200, gin.H{"content": "更新成功"})
}

func (s *AiChatService) GetSessionConversation(c *gin.Context) {
	var sessionReg models.ChatSession
	err := c.ShouldBindJSON(&sessionReg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	conversations := service.GetAiChatConversationDB(s.DB.DB, strconv.Itoa(sessionReg.SessionID))
	c.JSON(200, conversations)
}
