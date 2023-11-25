package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	InventoryUseCase services.InventoryUseCase
}

func NewInventoryHandler(usecase services.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{
		InventoryUseCase: usecase,
	}
}

func (i *InventoryHandler) AddInventory(c *gin.Context) {
	fmt.Println("*********** add inventory ***********")
	var inventory models.AddInventories

	if err := c.BindJSON(&inventory); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *InventoryHandler) ListProducts(c *gin.Context) {

	pageNo := c.DefaultQuery("page", "1")
	pageList := c.DefaultQuery("per_page", "5")
	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Product cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	pageListInt, err := strconv.Atoi(pageList)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Product cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	products_list, err := i.InventoryUseCase.ListProducts(pageNoInt, pageListInt)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Product cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	message := "product list"

	successRes := response.ClientResponse(http.StatusOK, message, products_list, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")
	err := i.InventoryUseCase.DeleteInventory(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the inventory", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *InventoryHandler) UpdateInventory(c *gin.Context) {

	var p models.InventoryUpdate

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := i.InventoryUseCase.UpdateInventory(p.Productid, p.Stock)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the inventory stock", a, nil)
	c.JSON(http.StatusOK, successRes)

}

// func (i *InventoryHandler) ListProductsForAdmin(c *gin.Context) {
// 	pageStr := c.Query("page")
// 	page, err := strconv.Atoi(pageStr)

// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// 	products, err := i.InventoryUseCase.ListProductsForAdmin(page)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
// 	c.JSON(http.StatusOK, successRes)
// }
