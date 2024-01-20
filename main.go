package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Result struct {
	Addition       int
	Subtraction    int
	Multiplication int
	Division       int
}

func Calculation(c echo.Context) error {
	// Get Query Params
	num1 := c.QueryParam("num1")
	a, err := strconv.Atoi(num1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	num2 := c.QueryParam("num2")
	b, err := strconv.Atoi(num2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result := Result{
		Addition:       a + b,
		Subtraction:    a - b,
		Multiplication: a * b,
		Division:       a / b,
	}
	return c.JSON(http.StatusOK, result)
}

var DB *gorm.DB

func main() {
	// New Echo Instance
	e := echo.New()

	// Initiate DB
	dsn := "root:root@tcp(localhost:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d
	// Get Method
	e.GET("/calculate", Calculation)
	// Web Server start
	e.Logger.Fatal(e.Start(":8080"))
}
