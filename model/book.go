package model

import "gorm.io/gorm"

type Book struct {
	ID     uint   `gorm:"primary_key"`
	Title  string `gorm:"type:varchar(100); not null"`
	Author string `gorm:"type:varchar(100); not null"`
}

func (Book) CrateBook(db *gorm.DB, title, author string) error {
	book := Book{
		Title:  title,
		Author: author,
	}
	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (Book) GetBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (Book) UpdateBook(db *gorm.DB, id uint, title, author string) error {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return err
	}
	book.Title = title
	book.Author = author
	if err := db.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func (Book) DeleteBook(db *gorm.DB, id uint) error {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}
