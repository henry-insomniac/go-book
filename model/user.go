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
func (User) CreateUser(db *gorm.DB, userName, email, password, phone string) (uint, error) {
	user := User{
		Username: userName,
		Email:    email,
		Password: password,
		Phone:    phone,
	}
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// UpdatePassword 修改密码
func (User) UpdatePassword(db *gorm.DB, id uint, password string) error {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}
	user.Password = password
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
