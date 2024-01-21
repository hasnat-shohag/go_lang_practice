package repositories

import (
	"demo/pkg/domain"
	"demo/pkg/models"
	"demo/pkg/types"
	"gorm.io/gorm"
)

// parent struct to implement interface binding
type bookRepo struct {
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

func (repo *bookRepo) GetBooks(bookID uint) []models.Book
func (repo *bookRepo) CreateBook(book *models.Book) error
func (repo *bookRepo) UpdateBook(book *models.Book) error
func (repo *bookRepo) DeleteBook(bookID uint) error

type BookService struct {
	repo domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		repo: bookRepo,
	}
}

func (service *BookService) GetBooks(bookID uint) ([]types.BookRequest, error)
func (service *BookService) CreateBook(book *models.Book) error
func (service *BookService) UpdateBook(book *models.Book) error
func (service *BookService) DeleteBook(bookID uint) error
