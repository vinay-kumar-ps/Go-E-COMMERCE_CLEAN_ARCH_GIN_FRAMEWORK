package interfaces

import "ecommerce/pkg/domain"

type CategoryUsecase interface{
	AddCategory (category string) (domain.Category,error)
	UpdateCategory (current ,new string)(domain.Category,error)
	DeleteCategory(CategoryId string)error
	GetCategories(page ,limit int) ([]domain.Category,error)
	
}