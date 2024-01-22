package containers

import (
	"demo/pkg/config"
	"demo/pkg/connection"
	"demo/pkg/controllers"
	"demo/pkg/repositories"
	"demo/pkg/routes"
	"demo/pkg/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	// config initialization
	config.SetConfig()
	// database connection
	connection.Connect()
	db := connection.GetDB()
	bookRepo := repositories.BookDBInstance(db)
	bookService := services.BookServiceInstance(bookRepo)
	bookCtr := controllers.NewBookController(bookService)
	b := routes.BookRoutes(e, bookCtr)

	b.InitBookRoute()

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
