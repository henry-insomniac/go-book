package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/service"
	"net/http"
)

type BookController struct {
	Service *service.BookService
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var input struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.Service.CreateBook(input.Name, input.Author); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}
