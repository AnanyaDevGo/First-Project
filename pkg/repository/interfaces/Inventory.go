package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type InventoryRepository interface {
	// AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	// CheckInventory(pid int) (bool, error)
	// UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	// DeleteInventory(id string) error
	// ShowIndividualProducts(id string) (models.Inventories, error)
	// ListProducts(int, int) ([]models.ProductsResponse, error)
	// ListProductsByCategory(id int) ([]models.Inventories, error)
	// CheckStock(inventory_id int) (int, error)
	// SearchProducts(key string) ([]models.Inventories, error)
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	ListProducts(int, int) ([]models.ProductsResponse, error)
	EditInventory(domain.Inventories, int) (domain.Inventories, error)
	DeleteInventory(id string) error
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	ShowIndividualProducts(id string) (models.ProductsResponse, error)
	CheckStock(inventory_id int) (int, error)
	FetchProductDetails(productId uint) (models.Inventories, error)
	GetInventory(prefix string) ([]models.ProductsResponse, error)
	FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error)
}
