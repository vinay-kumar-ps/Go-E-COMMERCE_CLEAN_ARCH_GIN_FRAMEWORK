package usecase

import (
	"ecommerce/pkg/helper"
	"ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"mime/multipart"
	"strconv"
)

type InventoryUsecase struct {
	invRepo interfaces.InventoryRespository
}

//contructor function

func NewInventoryUsecase(invRepo interfaces.InventoryRespository) services.InventoryUsecase {
	return InventoryUsecase{
		invRepo: invRepo,
	}
}
func (invU *InventoryUsecase) AddInventory(inventory models.Inventory, image *multipart.FileHeader) (models.InventoryResponse, error) {
	url, err := helper.AddImageToS3(image)

	if err != nil {
		return models.InventoryResponse{}, err
	}
	inventory.Image = url

	//send the url save in db
	inventoryResponse, err := invU.AddInventory(inventory, url)
	if err != nil {
		return models.InventoryResponse{}, err

	}
	return inventoryResponse, nil
}
func (invU *InventoryUsecase) UpdateImage(invID int, image *multipart.FileHeader) (models.Inventory, error) {

	url, err := helper.AddImageToS3(image)
	if err != nil {
		return models.Inventory{}, err
	}
	inventoryResponse, err := invU.UpdateImage(invID,url)
	if err != nil {
		return models.Inventory{}, err
	}
	return inventoryResponse, nil
}

func (invU *InventoryUsecase) UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error) {

	result, err := invU.invRepo.CheckInventory(invID)
	if err != nil {
		return models.Inventory{}, err
	}
	if !result {
		return models.Inventory{}, errors.New("there is no inventory as you mentioned")
	}
	newinventory,err :=invU.UpdateInventory(invID,invData)
	if err !=nil{
		return models.Inventory{},err
	}
	return newinventory,nil
}

func(invU *InventoryUsecase) DeleteInventory(id string)error{
	if err := invU.invRepo.DeleteInventory(id);err!=nil{
		return err
	}
	return nil
}
func (invU *InventoryUsecase)ShowIndividualProducts(id string)(models.InventoryDetails,error){
	product ,err :=invU.invRepo.ShowIndividualProducts(id)
	if err !=nil{
		return models.InventoryDetails{},err
	}
	productId,err:= strconv.Atoi(id)
	if err !=nil{
		return models.InventoryDetails{},err
	}
	var AdditionalImage []models.ImagesInfo
	AdditionalImage,err =invU.invRepo.GetImagesFromInventoryId(productId)
	if err !=nil{
		return models.InventoryDetails{},err
	}
    invDetails :=models.InventoryDetails{Inventory: product,AdditionalImages: AdditionalImage}
	return invDetails,nil
}

func (invU *InventoryUsecase) ListProducts(page int ,limit int)([]models.InventoryList,error){
	productDetails ,err :=invU.invRepo.ListProducts(page,limit)
	if err!=nil{
		return []models.InventoryList{},err
	}
	return productDetails,nil
}
func(invU *InventoryUsecase) GetCategoryProducts(catID int ,page,limit int)([]models.InventoryList,error){
	prdductDetails,err :=invU.invRepo.GetCategoryProducts(catID,page,limit)
	if err!= nil{
		return []models.InventoryList{},err
	}
	return prdductDetails,nil
}
func (invU *InventoryUsecase) AddImage(product_id int,image *multipart.FileHeader)(mode.InventoryResponse,error){
	//adding the to Aws s3 bucket
  imageUrl,err :=helper.AddImageToS3(image)
  if err !=nil{
	return models.InventoryResponse{},err
  }
  inventoryResponse,err :=invU.invRepo.AddImage(product_id,imageUrl)
  if err !=nil{
	return models.InventoryResponse{},err
  }
  return inventoryResponse,nil
}
func (invU *InventoryUsecase)DeleteImage(product_id ,image_id int)error{
	if err :=invU.invRepo.DeleteImage(product_id,image_id);err !=nil{
		return errors.New("image not deleted")
	}
	return nil
}
