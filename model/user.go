package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;type:varchar(32)" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Username  string    `gorm:"type:varchar(100);not null;unique"`
	Email     string    `gorm:"type:varchar(100);unique"`
	Phone     string    `gorm:"type:varchar(100);unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
}

// BeforeCreate 是 GORM 的 hook，在插入前自动执行
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID, err = gonanoid.New()
	}
	return
}

// CreateUser 创建用户
func (User) CreateUser(db *gorm.DB, userName, email, password, phone string) (string, error) {
	user := User{
		Username: userName,
		Email:    email,
		Password: password,
		Phone:    phone,
	}
	if err := db.Create(&user).Error; err != nil {
		return "", err
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
