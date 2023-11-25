package interfaces

import "CrocsClub/pkg/utils/models"

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	ListProducts(pageNo, pageList int) ([]models.ProductsResponse, error)
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	//ListProductsForAdmin(page int) ([]models.Inventories, error)
}
