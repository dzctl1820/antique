package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		// 验证Authorization格式（Bearer <token>）
		if len(token) < 7 {
			c.JSON(401, gin.H{"error": "无效的授权格式"})
			c.Abort()
			return
		}

		// 验证Token（调用VerifyJwtToken函数）
		claims, err := ParseJwt(token)
		log.Print(err)
		if err != nil {
			c.JSON(401, gin.H{"error": "无效或过期的Token"})
			c.Abort()
			return
		}

		// 将用户ID和角色存储到上下文（可选）
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

	}
}
