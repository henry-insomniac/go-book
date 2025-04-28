package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Username  string    `gorm:"type:varchar(100);not null;unique"`
	Email     string    `gorm:"type:varchar(100);unique"`
	Phone     string    `gorm:"type:varchar(100);unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
}

// CreateUser 创建用户
func (User) CreateUser(db *gorm.DB, userName, email, password, phone string) error {
	user := User{
		Username: userName,
		Email:    email,
		Password: password,
		Phone:    phone,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
