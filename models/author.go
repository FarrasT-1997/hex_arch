package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
