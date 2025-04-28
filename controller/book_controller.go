package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/service"
	"net/http"
	"strconv"
)

type BookController struct {
	Service *service.BookService
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var input struct {
		Name   string `json:"title"`
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Book created successfully", "status": 200})
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	books, err := c.Service.GetBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
	}
	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	var input struct {
		Name   string `json:"title"`
		Author string `json:"author"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	if err := c.Service.UpdateBook(uint(idUint), input.Name, input.Author); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "status": 200})
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	if err := c.Service.DeleteBook(uint(idUint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "status": 200})
}
