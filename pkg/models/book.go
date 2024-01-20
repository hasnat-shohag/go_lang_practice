package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName    string
	Author      string
	Publication string
}
