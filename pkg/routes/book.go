package routes

import (
	"demo/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	book := e.Group("/bookstore")
	book.POST("/book", controllers.CreateBook)
	book.GET("/book", controllers.GetBook)
	book.PUT("/book", controllers.UpdateBook)
	book.DELETE("/book/:bookID", controllers.DeleteBook)
}
