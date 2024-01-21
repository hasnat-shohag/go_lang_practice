package controllers

import (
	"demo/pkg/domain"
	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {}
func GetBook(c echo.Context) error    {}
func UpdateBook(c echo.Context) error {}
func DeleteBook(c echo.Context) error {}

var BookService domain.IBookService

func SetBookService(bService domain.IBookService) {
	BookService = bService
}
