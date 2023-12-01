package usecase

import (
	"CrocsClub/pkg/domain"
	helper_interface "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	usecase "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"mime/multipart"

	"strings"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
	helper     helper_interface.Helper
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interface.Helper) usecase.InventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
		helper:     h,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.AddInventories, file *multipart.FileHeader) (models.ProductsResponse, error) {

	if inventory.Stock < 0 || inventory.Price < 0 || inventory.CategoryID < 0 {
		err := errors.New("enter valid values")
		return models.ProductsResponse{}, err
	}

	url, err := i.helper.AddImageToAwsS3(file)
	if err != nil {
		return models.ProductsResponse{}, err
	}
	productResponse, err := i.repository.AddInventory(inventory, url)
	if err != nil {
		return models.ProductsResponse{}, err
	}
	return productResponse, err
}

func (i *inventoryUseCase) ListProducts(pageNo, pageList int) ([]models.ProductsResponse, error) {

	offset := (pageNo - 1) * pageList
	productList, err := i.repository.ListProducts(pageList, offset)
	if err != nil {
		return []models.ProductsResponse{}, err
	}
	return productList, nil
}

func (usecase *inventoryUseCase) EditInventory(inventory domain.Inventories, id int) (domain.Inventories, error) {
	modInventory, err := usecase.repository.EditInventory(inventory, id)
	if err != nil {
		return domain.Inventories{}, err
	}
	return modInventory, nil
}

func (usecase *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := usecase.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil
}

func (i inventoryUseCase) UpdateInventory(pid int, stock int) (models.ProductsResponse, error) {

	result, err := i.repository.CheckInventory(pid)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	if !result {
		return models.ProductsResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	return newcat, err
}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.ProductsResponse, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	return product, nil
}

func (i *inventoryUseCase) SearchProductsOnPrefix(prefix string) ([]models.ProductsResponse, error) {

	inventoryList, err := i.repository.GetInventory(prefix)

	if err != nil {
		return nil, err
	}

	var filteredProducts []models.ProductsResponse

	for _, product := range inventoryList {
		if strings.HasPrefix(strings.ToLower(product.ProductName), strings.ToLower(prefix)) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if len(filteredProducts) == 0 {
		return nil, errors.New("no items matching your keyword")
	}

	return filteredProducts, nil
}

func (i *inventoryUseCase) FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error) {
	product_list, err := i.repository.FilterByCategory(CategoryIdInt)

	if err != nil {
		return []models.ProductsResponse{}, err
	}

	return product_list, nil
}
