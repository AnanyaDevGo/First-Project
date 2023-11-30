package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	ListProducts(int, int) ([]models.ProductsResponse, error)
	EditInventory(domain.Inventories, int) (domain.Inventories, error)
	DeleteInventory(id string) error
	UpdateInventory(productID int, stock int) (models.InventoryResponse, error)
	ShowIndividualProducts(id string) (models.ProductsResponse, error)
	SearchProductsOnPrefix(prefix string) ([]models.ProductsResponse, error)
	FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error)
}
