package models

import "gorm.io/gorm"

type BookDetail struct {
	gorm.Model
	BookName    string
	Author      string
	Publication string
}
