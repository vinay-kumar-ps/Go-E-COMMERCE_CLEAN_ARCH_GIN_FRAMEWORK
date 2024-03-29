package interfaces

import (
	"ecommerce/pkg/domain"

)
type CategoryUsecase interface {
	AddCategory(category string) (domain.Category, error)
	UpdateCategory(currrent, new string) (domain.Category, error)
	DeleteCategory(categoryId string) error
	GetCategories() ([]domain.Category, error)
}
