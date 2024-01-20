package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

var DB *gorm.DB

func CreateUsers(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "User created successfully")
}
func GetAllUsers(c echo.Context) error {
	users := []User{}
	if err := DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
func DeleteUsers(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user := &User{}
	if err := DB.Where("id = ?", ID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "user Deleted successfully")
}

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
	DB.AutoMigrate(&User{})

	e.GET("/users", GetAllUsers)
	e.POST("/users", CreateUsers)
	e.DELETE("/users/:id", DeleteUsers)

	e.Logger.Fatal(e.Start(":1111"))
}
