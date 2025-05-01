package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/henry-insomniac/go-book/service"
	"net/http"
	"strconv"
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

	userID, err := c.Service.CreateUser(user.Username, user.Email, user.Phone, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userResponse := struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}{
		ID:       userID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}

// ForgetPassword 忘记密码
func (c *UserController) ForgetPassword(ctx *gin.Context) {
	var input struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	// 使用 ShouldBindJSON 来绑定请求体中的 JSON 数据
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// string 转 unit
	// 验证 id 是否为有效的数字字符串
	idUint, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		// 如果 id 无法转换为数字，返回错误
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, must be a number"})
		return
	}

	if err := c.Service.UpdatePassword(uint(idUint), input.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
}
