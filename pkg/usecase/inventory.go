package usecase

import (
	"ecommerce/pkg/helper"
	"ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"mime/multipart"
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
func (invU *InventoryUsecase)AddInventory(inventory models.Inventory,image *multipart.FileHeader)(models.InventoryResponse,error){
	url ,err :=helper.AddImageToS3(image)
 
	if err !=nil{
		return models.InventoryResponse{},err
	}
	inventory.Image =url
	//send the url save in db
	inventoryResponse,err :=invU.AddInventory(inventory,url)
	if err!=nil{
		return models.InventoryResponse{},err

	}
	return inventoryResponse,nil
}
func (invU *InventoryUsecase)UpdateImage(invID int ,image *multipart.FileHeader)(models.Inventory,error){

	url,err :=helper.AddImageToS3(image)
	if err !=nil{
		return models.Inventory{},err
	}
	inve
}

