package database

import "gorm.io/gorm"

type DatabaseService struct {
	DB *gorm.DB
}
