package connection

import (
	"demo/pkg/config"
	"demo/pkg/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIP, dbConfig.DBName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

func migrate() {
	db.Migrator().AutoMigrate(&models.Book{})
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
