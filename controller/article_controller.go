package controller

import (
	"github.com/henry-insomniac/go-book/model"
	"github.com/henry-insomniac/go-book/service"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	Service *service.ArticleService
}

// CreateArticle POST /articles
func (ac *ArticleController) CreateArticle(ctx *gin.Context) {
	var input struct {
		Title       string   `json:"title"`
		ContentMD   string   `json:"content_md"`
		Summary     string   `json:"summary"`
		CoverImage  string   `json:"cover_image"`
		CoverThumb  string   `json:"cover_thumb"`
		ReadTime    int      `json:"read_time"`
		PublishedAt string   `json:"published_at"`
		AuthorID    string   `json:"author_id"`
		TagIDs      []string `json:"tag_ids"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	publishTime, err := time.Parse("2006-01-02", input.PublishedAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid published_at format"})
		return
	}

	article := model.Article{
		Title:       input.Title,
		ContentMD:   input.ContentMD,
		Summary:     input.Summary,
		CoverImage:  input.CoverImage,
		CoverThumb:  input.CoverThumb,
		ReadTime:    input.ReadTime,
		PublishedAt: publishTime,
		AuthorID:    input.AuthorID,
	}

	// 加载标签
	for _, tagID := range input.TagIDs {
		article.Tags = append(article.Tags, model.Tag{ID: tagID})
	}

	if err := ac.Service.CreateArticle(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article created", "id": article.ID})
}

// GetAllArticles GET /articles
func (ac *ArticleController) GetAllArticles(ctx *gin.Context) {
	articles, err := ac.Service.GetAllArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"articles": articles})
}

// SearchArticles GET /articles/search?keyword=xxx
func (ac *ArticleController) SearchArticles(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	articles, err := ac.Service.SearchArticles(keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search articles"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"articles": articles})
}
