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
	catrepo    interfaces.CategoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interface.Helper, catrepo interfaces.CategoryRepository) usecase.InventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
		helper:     h,
		catrepo:    catrepo,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.AddInventories, file *multipart.FileHeader) (models.ProductsResponse, error) {

	if inventory.Stock <= 0 || inventory.Price <= 0 || inventory.CategoryID <= 0 {
		err := errors.New("enter valid values")
		return models.ProductsResponse{}, err
	}
	if ok, _ := i.catrepo.CategoryExists(inventory.CategoryID); !ok {
		return models.ProductsResponse{}, errors.New("category does not exist")
	}

	if ok, _ := i.repository.CheckInventoryByCatAndName(inventory.CategoryID, inventory.ProductName); ok {
		return models.ProductsResponse{}, errors.New("already added")
	}
	productname, err := i.helper.ValidateAlphabets(inventory.ProductName)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	if !productname {
		return models.ProductsResponse{}, errors.New("invalid format for name")
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

func (i *inventoryUseCase) MultipleImageUploader(inventoryID int, files []*multipart.FileHeader) error {
	for _, file := range files {
		url, err := i.helper.AddImageToAwsS3(file)
		if err != nil {
			return err
		}
		err = i.repository.ImageUploader(inventoryID, url)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (i *inventoryUseCase) ListProducts(pageNo, pageList int) ([]models.ProductsResponse, error) {

// 	offset := (pageNo - 1) * pageList
// 	productList, err := i.repository.ListProducts(pageList, offset)
// 	if err != nil {
// 		return []models.ProductsResponse{}, err
// 	}
// 	return productList, nil
// }

func (i *inventoryUseCase) ListProducts(pageNo, pageList int) ([]models.ProductsResponse, error) {
	offset := (pageNo - 1) * pageList
	productList, err := i.repository.ListProducts(pageList, offset)
	if err != nil {
		return []models.ProductsResponse{}, err
	}

	// Assuming productList is of type []models.ProductsResponseDisp
	var convertedProductList []models.ProductsResponse
	for _, product := range productList {
		ConvertedProduct := models.ProductsResponse{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			ProductName: product.ProductName,
			Size:        product.Size,
			Price:       product.Price,
			Stock:       product.Stock,
			Image:       product.Image,
		}
		ConvertedProduct = models.ProductsResponse(product)
		convertedProductList = append(convertedProductList, ConvertedProduct)
	}

	return convertedProductList, nil
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

func (i inventoryUseCase) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {
	var res models.InventoryResponse
	if stock <= 0 {
		return models.InventoryResponse{}, errors.New("invalid input")
	}

	result, err := i.repository.CheckInventory(pid)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	if !result {
		return models.InventoryResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	res.ProductID = int(newcat.ID)
	res.Stock = newcat.Stock

	return res, err
}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.ProductsResponse, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.ProductsResponse{}, err
	}
	// DiscountPercentage, err := i.offerRepository.FindDiscountPercentage(product.CategoryID)
	// if err != nil {
	// 	return models.Inventories{}, err
	// }

	// //make discounted price by calculation
	// var discount float64
	// if DiscountPercentage > 0 {
	// 	discount = (product.Price * float64(DiscountPercentage)) / 100
	// }

	// product.DiscountedPrice = product.Price - discount

	return product, nil
}

func (i *inventoryUseCase) SearchProductsOnPrefix(prefix string) ([]models.ProductsResponse, error) {
	if prefix == "" {
		return []models.ProductsResponse{}, errors.New("name should not be empty")
	}

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
		return nil, errors.New("no items are matching")
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
