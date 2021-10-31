package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	Genre       string
	Year        int
}
