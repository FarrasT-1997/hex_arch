package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string
	Sex      string
	DoB      string
	Email    string
	Password string
	Token    string
}
