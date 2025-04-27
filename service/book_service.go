package service

import (
	"github.com/henry-insomniac/go-book/model"
	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

func (s *BookService) CreateBook(title, author string) error {
	return model.Book{}.CrateBook(s.DB, title, author)
}

func (s *BookService) UpdateBook(id uint, title, author string) error {
	return model.Book{}.UpdateBook(s.DB, id, title, author)
}

func (s *BookService) DeleteBook(id uint) error {
	return model.Book{}.DeleteBook(s.DB, id)
}

func (s *BookService) GetBook() ([]model.Book, error) {
	return model.Book{}.GetBooks(s.DB)
}
