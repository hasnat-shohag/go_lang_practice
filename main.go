package main

import (
	"demo/pkg/containers"
	"github.com/labstack/echo/v4"
)

func main() {
	// New Echo Instance
	e := echo.New()
	containers.Serve(e)
}
