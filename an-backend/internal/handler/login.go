package handler

import (
	"an-backend/internal/service"
	"an-backend/models"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	DB *service.DataDB
}

func (s *UserService) LoginByEmail(c *gin.Context) {
	var req models.LoginRequestByEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := s.DB.LoginByEmail(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
