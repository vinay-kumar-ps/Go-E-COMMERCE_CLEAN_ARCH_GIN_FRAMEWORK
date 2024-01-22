package usecase

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"errors"
)

type categoryUsecase struct{
	repo interfaces.CategoryRepository

}

//constructor function
func NewCategoryUsecase(repo interfaces.CategoryRepository) services.CategoryUsecase{
	return &categoryUsecase{
		repo: repo,
	}

}
func(catU *categoryUsecase) AddCategory(category string)(domain.Category,error){
	productResponse ,err :=catU.repo.AddCategory(category)
	if err != nil{
		return domain.Category{},err
	}
	return productResponse,nil
}
func (catU *categoryUsecase)UpdateCategory (current,new string)(domain.Category,error){
	result,err :=catU.repo.CheckCategory(current)
	if err !=nil{
		return domain.Category{},err
	}
	if !result{
		return domain.Category{},errors.New("no category as you mentioned")
	}
	newCat,err :=catU.UpdateCategory(current,new)
	if err !=nil{
		return domain.Category{},err
	}
	return newCat,nil
}
func (catU *categoryUsecase)DeleteCategory(categoryId string)error{
	err :=catU.repo.DeleteCategory(categoryId)
	if err !=nil{
		return err 
	}
	return nil
}
func (catU *categoryUsecase)GetCategories(page ,limit int) ([]domain.Category,error){
categories,err :=catU.repo.GetCategories(page,limit)
	if err !=nil{
	return []domain.Category{},errors.New("categories not found")

}
return categories,nil
}