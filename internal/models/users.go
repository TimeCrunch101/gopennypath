package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username" gorm:"uniqueIndex;size:45"`
	Password  string `json:"password"`
}
