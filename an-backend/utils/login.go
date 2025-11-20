package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"strings"
	"time"
)

type CustomClaims struct {
	Email                string `json:"email"`
	UserID               uint   `json:"user_id"`        // 添加用户ID
	Role                 string `json:"role,omitempty"` // 可选角色字段
	jwt.RegisteredClaims        // 嵌入标准声明
}

func GenerateJwtToken(userID uint, email string, expiresIn time.Duration) (string, error) {
	// 1. 创建自定义Claims（包含标准声明）
	claims := CustomClaims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // 签发时间
			Issuer:    "your_app_name",                               // 签发者
			Subject:   "user_auth",                                   // 主题
		},
	}

	// 2. 创建Token（推荐HS512或RS256算法）
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// 3. 从环境变量读取密钥（增强安全性）
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 64 { // 密钥长度检查
		return "", errors.New("JWT_SECRET must be at least 32 characters")
	}

	// 4. 签名并返回
	return token.SignedString([]byte(secret))
}

func ParseJwt(tokenString string) (*CustomClaims, error) {
	// 1. 检查token是否为空
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	// 2. 从环境变量读取密钥
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 32 {
		return nil, errors.New("JWT_SECRET must be at least 32 characters")
	}

	// 3. 解析Token
	token, err := jwt.ParseWithClaims(
		strings.TrimSpace(tokenString), // 去除前后空格
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法是否匹配
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("token parse error: %w", err)
	}

	// 4. 验证Token是否有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

func GenerateCode() string {
	rand.NewSource(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendMail(from string, to string, subject string, body string, authorizeCode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)       // 发送人
	m.SetHeader("To", to)           //  接收人
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/plain", body)   // 正文的内容。text/plain表示纯文本，"text/html" 表示 HTML 内容。
	d := gomail.NewDialer("smtp.qq.com", 587, from, authorizeCode)
	// 通过拨号器对象发送指定的邮件消息
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	// 注意：不再在这里存储验证码到Redis，因为这会在handler中完成
	// 避免重复存储导致的覆盖问题
	return nil
}
