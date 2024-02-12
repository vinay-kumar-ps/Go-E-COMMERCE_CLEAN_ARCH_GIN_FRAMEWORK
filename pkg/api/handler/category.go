package handler

import (
	services "ecommerce/pkg/usecase"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase

}
func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: usecase,
	}
}

