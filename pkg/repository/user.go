package repository

import "gorm.io/gorm"

type userRepository struct{
	DB *gorm.DB
}