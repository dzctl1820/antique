package handler

import (
	"an-backend/internal/service"
	"an-backend/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleService struct {
	DB *service.ArticleDB
}

func (s *ArticleService) AddArticle(c *gin.Context) {
	var article models.ArticleRequest
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 调用服务层创建文章
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(400, gin.H{"error": "userID not found"})
		return
	}
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(400, gin.H{"error": "userID type error"})
		return
	}
	article.IDUser = int(uid)
	if err := s.DB.CreateArticle(&article); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 返回创建成功
	c.JSON(200, gin.H{"message": "文章创建成功"})
}

func (s *ArticleService) GetArticle(c *gin.Context) {
	// 获取文章ID
	id, _ := strconv.Atoi(c.Param("id"))
	// 调用服务层获取文章
	article, err := s.DB.GetArticleById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	// 返回文章
	c.JSON(200, article)
}

func (s *ArticleService) GetArticles(c *gin.Context) {
	// 调用服务层获取所有文章
	articles, err := s.DB.GetAllArticles()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 返回所有文章
	c.JSON(200, articles)
}
