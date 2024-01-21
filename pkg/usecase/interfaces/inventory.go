package interfaces

import (
	"ecommerce/pkg/utils/models"
	"mime/multipart"
)

type InventoryUsecase interface{
	AddInventory (inventory models.Inventory,image *multipart.FileHeader)(models.InventoryResponse,error)
	UpdateInventory (invID int, invDate models.UpdateInventory)(models.Inventory,error)
	UpdateImage(invID int ,image *multipart.FileHeader)(models.Inventory,error)
	DeleteInventory (id string)error

	ShowIndividualProducts(id string) (models.InventoryDetails,error)
	ListProduts(page int, limit int)([]models.InventoryList,error)
	SearchProducts(key string,page ,limit int)([]models.InventoryList,error)
	GetCategoryProducts(cartID int ,page,limit int )([]models.InventoryList,error)
	AddImage(product_id int ,image *multipart.FileHeader)(models.InventoryResponse,error)
	DeleteImage(product_id,image_id int)error
}