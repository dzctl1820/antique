package service

import (
	"an-backend/models"
	"gorm.io/gorm"
)

type ArticleDB struct {
	DB *gorm.DB
}

func (s *ArticleDB) CreateArticle(articleRequest *models.ArticleRequest) error {
	article := &models.Article{
		Author:    articleRequest.Author,
		IDUser:    articleRequest.IDUser,
		Cover:     articleRequest.Cover,
		Category:  articleRequest.Category,
		Content:   articleRequest.Content,
		Summary:   articleRequest.Summary,
		Tags:      articleRequest.Tags,
		Title:     articleRequest.Title,
		Status:    "published",
		Favorites: 0,
		Likes:     0,
		Views:     0,
	}
	return s.DB.Create(article).Error
}

func (s *ArticleDB) GetArticleById(id int) (*models.Article, error) {
	var article models.Article
	err := s.DB.Where("id = ?", id).First(&article).Error
	return &article, err
}

func (s *ArticleDB) GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	err := s.DB.Find(&articles).Error
	return articles, err
}
