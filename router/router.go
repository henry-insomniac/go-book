package router

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/controller"
	"github.com/henry-insomniac/go-book/database"
	"github.com/henry-insomniac/go-book/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create a BookService instance
	bookService := &service.BookService{
		DB: database.DB, // Assume database.DB is your gorm.DB instance
	}

	// Create a BookController instance
	bookController := &controller.BookController{
		Service: bookService,
	}

	// Define routes
	r.POST("/books", bookController.CreateBook)
	r.GET("/books", bookController.GetBooks)
	r.PUT("/books/:id", bookController.UpdateBook)
	r.DELETE("/books/:id", bookController.DeleteBook)

	return r
}
