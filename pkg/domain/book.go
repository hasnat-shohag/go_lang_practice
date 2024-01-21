package domain

import (
	"demo/pkg/models"
	"go/types"
)

// Database Interface
type IBookRepo interface {
	GetBooks(bookID uint) []models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}

// Service Interface
type IBookService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}
