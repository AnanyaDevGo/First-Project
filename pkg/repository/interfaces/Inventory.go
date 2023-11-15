package interfaces

import "CrocsClub/pkg/utils/models"

type InventoryRepository interface {
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (models.Inventories, error)
	ListProducts(page int) ([]models.Inventories, error)
	ListProductsByCategory(id int) ([]models.Inventories, error)
	CheckStock(inventory_id int) (int, error)
	SearchProducts(key string) ([]models.Inventories, error)
}
