package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)
 
type categoryRepository struct{
	DB *gorm.DB
}
func NewCategoryRepository(db *gorm.DB)interfaces.CategoryRepository  {
	return &categoryRepository{
		DB: db,
	}
	
}
func (cat *cartRepository) AddCategory(category string) (domain.Category,error){
	var b string
}