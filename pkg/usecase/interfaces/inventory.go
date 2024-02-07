package interfaces

import (
	"ecommerce/pkg/utils/models"
	"mime/multipart"
)


type InventoryUseCase interface {
	AddInventory(inventory models.AddInventories, image *multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(sku string) (models.Inventories, error)
	ListProductsForUser(page, userID int) ([]models.Inventories, error)
	ListProductsForAdmin(page int) ([]models.Inventories, error)

	SearchProducts(key string) ([]models.Inventories, error)

	UpdateProductImage(id int, file *multipart.FileHeader) error
	EditInventoryDetails(int, models.EditInventoryDetails) error
}