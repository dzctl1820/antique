package handler

import (
	"an-backend/database"
	"an-backend/internal/service"
	"an-backend/models"
	"an-backend/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type UserService struct {
	DB *service.UserDB
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

func (s *UserService) SendCodeByEmail(c *gin.Context) {
	var req models.SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 生成验证码
	code := utils.GenerateCode()

	err := utils.SendMail("3186807685@qq.com", req.Email, "验证码", code, "ktplhxrwssonddgf")
	if err != nil {
		log.Printf("发送邮件失败: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 将验证码存储到Redis中
	log.Printf("准备存储验证码到Redis: email=%s, code=%s", req.Email, code)
	err = database.InsertRedis(req.Email+"Code", code)
	if err != nil {
		log.Printf("存储验证码到Redis失败: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Printf("验证码存储成功: email=%s, code=%s", req.Email, code)
	c.JSON(200, gin.H{"message": "发送成功"})
}

func (s *UserService) RegisterByEmail(c *gin.Context) {
	var req models.RegisterRequestByEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 从Redis中获取验证码
	code, err := database.FindRedis(req.Email + "Code")
	if err != nil {
		c.JSON(400, gin.H{"error": "验证码已过期或不存在"})
		return
	}

	if code != req.Code {
		c.JSON(400, gin.H{"error": "验证码错误"})
		return
	}
	err = s.DB.RegisterByEmail(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "注册成功"})
}
