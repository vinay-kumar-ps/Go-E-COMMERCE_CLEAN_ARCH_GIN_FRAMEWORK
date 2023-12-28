package repository

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

//constructor function

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {

	return &inventoryRepository{
		DB: DB,
	}
}

func (ir *inventoryRepository) AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error) {
	var inventoryResp models.InventoryResponse

	query := `INSERT INTO inventories (category_id,product_name,description,stock,price,image)
	VALUES(?,?,?,?,?,?)RETURNING id`

	err := ir.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price, inventory.Image, url).Scan(&inventoryResp.ProductID).Error
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return models.InventoryResponse{}, nil
}

func (ir *inventoryRepository) UpdateImage(inventId int, url string) (models.Inventory, error) {

	//check db connecction
	if ir.DB == nil {
		return models.Inventory{}, errors.New("database connection failed while updating image")

	}

	//updating image
	err := ir.DB.Exec("UPDATEN inventories SET image =? WHERE id =?",inventId).Scan(&ir.UpdateImageInventory).Error

  if err !=nil{
	return models.Inventory{},err
  }
  //Retrive the update
  var UpdateImageInventory models.Inventory
  err
}
