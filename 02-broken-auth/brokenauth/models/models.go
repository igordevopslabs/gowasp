package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Token struct {
	gorm.Model
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
}
