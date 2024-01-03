package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
	"mime/multipart"
)

type InventoryUseCase interface {
	AddInventory(product models.AddInventories, file *multipart.FileHeader) (models.ProductsResponse, error)
	ListProducts(int, int) ([]models.ProductsResponse, error)
	EditInventory(domain.Inventories, int) (domain.Inventories, error)
	DeleteInventory(id string) error
	UpdateInventory(productID int, stock int) (models.InventoryResponse, error)
	ShowIndividualProducts(id string) (models.ProductsResponse, error)
	SearchProductsOnPrefix(prefix string) ([]models.ProductsResponse, error)
	FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error)
	MultipleImageUploader( inventoryID int,files []*multipart.FileHeader) error
}
