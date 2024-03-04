package usecase

import (
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"errors"
	"fmt"
)

type categoryUsecase struct {
	repo interfaces.CategoryRepository
}

// constructor function

func NewCategoryUsecase(repo interfaces.CategoryRepository) services.CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

func (catU *categoryUsecase) AddCategory(category string) (domain.Category, error) {
	productResponse, err := catU.repo.AddCategory(category)
	if err != nil {
		fmt.Println("error from cat usecase ", err)
		return domain.Category{}, err
	}
	return productResponse, nil
}

func (catU *categoryUsecase) UpdateCategory(currrent, new string) (domain.Category, error) {
	result, err := catU.repo.CheckCategory(currrent)
	if err != nil {
		return domain.Category{}, err
	}
	if !result {
		return domain.Category{}, errors.New("no category as you mentioned")
	}
	newCat, err := catU.UpdateCategory(currrent, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCat, nil
}

func (catU *categoryUsecase) DeleteCategory(categoryId string) error {
	err := catU.repo.DeleteCategory(categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (catU *categoryUsecase) GetCategories() ([]domain.Category, error) {
	categories, err := catU.repo.GetCategories()
	if err != nil {
		return []domain.Category{}, errors.New("categories not found")
	}
	return categories, nil
}
