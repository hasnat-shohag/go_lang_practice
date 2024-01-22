package domain

import (
	"demo/pkg/models"
	"demo/pkg/types"
)

// Database Interface
type IBookRepo interface {
	GetBooks(bookID uint) []models.BookDetail
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
}

// Service Interface
type IBookService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
}
