package service

import (
	"github.com/henry-insomniac/go-book/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUser(userName, email, phone, password string) error {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return model.User{}.CreateUser(s.DB, userName, email, string(hash), phone)
}
