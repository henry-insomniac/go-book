package service

import (
	"github.com/henry-insomniac/go-book/model"
	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

// CreateArticleService 创建文章
func (s *ArticleService) CreateArticle(article *model.Article) error {
	return s.DB.Create(article).Error
}

// GetAllArticles 获取所有文章（可分页）
func (s *ArticleService) GetAllArticles() ([]model.Article, error) {
	var articles []model.Article
	err := s.DB.Preload("Author").Preload("Tags").Find(&articles).Error
	return articles, err
}

// SearchArticles 搜索
func (s *ArticleService) SearchArticles(keyword string) ([]model.Article, error) {
	var articles []model.Article
	err := s.DB.Preload("Author").Preload("Tags").
		Where("title LIKE ? OR summary LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Find(&articles).Error
	return articles, err
}
