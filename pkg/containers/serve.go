package containers

import (
	"demo/pkg/config"
	"demo/pkg/connection"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	// config initialization
	config.SetConfig()
	// database connection
	connection.Connect()

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
