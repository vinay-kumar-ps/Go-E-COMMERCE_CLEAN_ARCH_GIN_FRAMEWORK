package repository

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"strconv"

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

func (ir *inventoryRepository) UpdatedImage(inventId int, url string) (models.Inventory, error) {

	//check db connecction
	if ir.DB == nil {
		return models.Inventory{}, errors.New("database connection failed while updating image")

	}

	//updating image
	err := ir.DB.Exec("UPDATEN inventories SET image =? WHERE id =?",inventId).Scan(&ir.UpdatedImageInventory).Error

  if err !=nil{
	return models.Inventory{},err
  }
  //Retrive the update
  var UpdatedImageInventory models.Inventory
  err= ir.DB.Raw("SELECT COUNT  * FROM  inventories WHERE id =?",inventId).Scan(&UpdatedImageInventory).Error
   
  if err!=nil{
	return models.Inventory{},err

  }
  return UpdatedImageInventory,nil
}
func (ir *inventoryRepository) CheckInventory(pid int) (bool,error) {

	var check int
	err := ir.DB.Raw("SELECT COUNT (*) FROM inventories WHERE id =?",pid).Scan(&check).Error
	 
	if err!=nil{
return false ,err
	}
	if check ==0{
		return false,err
	}
	return true,nil
}

func (ir *inventoryRepository) UpdateInventory(pid int,invData models.UpdateInventory)(models.Inventory,error) {
	if  ir.DB==nil{
		return models.Inventory{},errors.New("databse connection failed while update inventory")
	}


	if invData.CategoryID !=0 {
		if err := ir.DB.Exec("UPDATE inventories SET category_id=?,WHERE id=?",invData.CategoryID,pid).Error;err !=nil {

        return models.Inventory{},err
		}


	}

if invData.ProductName !="" && invData.ProductName!= "string" {
		if err := ir.DB.Exec("UPDATE inventories SET product_name=?,WHERE id=?",invData.ProductName,pid).Error; err!=nil {

        return models.Inventory{},err
		}

	}


	if invData.Description !="" && invData.Description!= "string" {
		if err := ir.DB.Exec("UPDATE inventories SET description=?,WHERE id=?",invData.Description,pid).Error; err!=nil {

        return models.Inventory{},err
		}

	}
	

   if invData.Stock !=0 {
	if err := ir.DB.Exec("UPDATE inventories SET stock =? WHERE id =?",invData.Stock,pid).Error;err!=nil{

		return models.Inventory{},err
	}

 
}

if invData.Price !=0 {
	if err := ir.DB.Exec("UPDATE inventories SET price =? WHERE id=?",invData.Price.pid).Error;err!=nil{

		return models.Inventory{},err
	}
}

	//retrive the updates
	 var updatedInventory models.Inventory
	 err := ir.DB.Raw("SELECT * FROM inventories WHERE id =?",pid).Scan(&updatedInventory).Error
     if err!=nil{
		return models.Inventory{},err
	 }

     return updatedInventory,nil
}

func(ir *inventoryRepository) DeleteInventory(inventoryId string)error{
	id,err := strconv.Atoi(inventoryId)

	if err!= nil{
		return errors.New("string to int conversion failed")
	}

	result :=ir.DB.Exec("DELETE FROM inventories WHERE id =?",id)

	if result.RowsAffected <1 {
		return errors.New("no records exists with this id")
	}
	return nil

}
func