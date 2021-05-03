package models

import "github.com/jinzhu/gorm"

// User model => Describes the user model
type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Base
}
