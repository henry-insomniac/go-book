package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/service"
	"net/http"
)

type UserController struct {
	Service *service.UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	if err := c.Service.CreateUser(user.Username, user.Email, user.Phone, user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userResponse := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}
