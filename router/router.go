package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/controller"
	"github.com/henry-insomniac/go-book/database"
	"github.com/henry-insomniac/go-book/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 或指定具体的域名，比如 http://localhost:3000
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Create a BookService instance
	bookService := &service.BookService{
		DB: database.DB, // Assume database.DB is your gorm.DB instance
	}

	userService := &service.UserService{
		DB: database.DB,
	}

	articleService := &service.ArticleService{
		DB: database.DB,
	}

	// Create a BookController instance
	bookController := &controller.BookController{
		Service: bookService,
	}

	userController := &controller.UserController{
		Service: userService,
	}

	articleController := &controller.ArticleController{
		Service: articleService,
	}

	// Define routes
	r.POST("/interface/books", bookController.CreateBook)
	r.GET("/interface/books", bookController.GetBooks)
	r.PUT("/interface/books/:id", bookController.UpdateBook)
	r.DELETE("/interface/books/:id", bookController.DeleteBook)
	r.POST("/interface/createUser", userController.CreateUser)
	r.POST("/interface/forgetPassword", userController.ForgetPassword)

	// 博客路由
	r.POST("/interface/articles", articleController.CreateArticle)
	r.GET("/interface/articles", articleController.GetAllArticles)
	r.GET("/interface/articles/search", articleController.SearchArticles)
	r.GET("/interface/articles/:id", articleController.GetArticleByID)

	return r
}
