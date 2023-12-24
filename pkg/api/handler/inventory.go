package handler

import (
	"CrocsClub/pkg/domain"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
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

// @Summary Add Inventory
// @Description Add a new inventory item.
// @Accept multipart/form-data
// @Produce json
// @security BearerTokenAuth
// @Tags Admin Product Management
// @Param category_id formData int true "ID of the category for the inventory"
// @Param product_name formData string true "Name of the product"
// @Param size formData string true "Size of the product"
// @Param stock formData int true "Stock quantity of the product"
// @Param price formData float64 true "Price of the product"
// @Param image formData file true "Image file of the product"
// @Success 200 {object} response.Response "Successfully added Product"
// @Failure 400 {object} response.Response "Invalid request format or fields provided in the wrong format"
// @Failure 500 {object} response.Response "Failed to add the product"
// @Router /admin/inventories [post]
func (i *InventoryHandler) AddInventory(c *gin.Context) {
	var inventory models.AddInventories

	cat := c.PostForm("category_id")
	inventory.CategoryID, _ = strconv.Atoi(cat)
	inventory.ProductName = c.PostForm("product_name")
	inventory.Size = c.PostForm("size")
	inventory.Stock, _ = strconv.Atoi(c.PostForm("stock"))
	inventory.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	file, err := c.FormFile("image")
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from the Form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productResponse, err := i.InventoryUseCase.AddInventory(inventory, file)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Product", productResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// MultipleImageUploader uploads multiple images for a specific inventory item.
//
// @Summary Upload multiple images for an inventory item
// @Description Upload multiple images for a specific inventory item using the provided inventory ID and images.
// @Tags Admin Product Management
// @security BearerTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param inventory_id formData integer true "Inventory ID for which images are uploaded"
// @Param image formData file true "Images to be uploaded" collection
// @Success 200 {object} response.Response "Successfully uploaded images"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 502 {object} response.Response "Bad Gateway"
// @Router /admin/inventories/uploadimages [post]
func (i *InventoryHandler) MultipleImageUploader(c *gin.Context) {
	inventoryID := c.PostForm("inventory_id")
	inventoryIDint, err := strconv.Atoi(inventoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "retrieving form data error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	files := form.File["image"]
	if len(files) == 0 {
		errRes := response.ClientResponse(http.StatusBadRequest, "no images provided", nil, "")
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = i.InventoryUseCase.MultipleImageUploader(inventoryIDint, files)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "uploading failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully uploaded images", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary List Products
// @Description Get a paginated list of products.
// @Accept json
// @Produce json
// @Tags Admin Product Management
// @Param page query int false "Page number for pagination (default: 1)"
// @Param per_page query int false "Number of products per page (default: 5)"
// @Success 200 {object} response.Response "Product list retrieved successfully"
// @Failure 400 {object} response.Response "Invalid request format or fields provided in the wrong format"
// @Failure 500 {object} response.Response "Failed to retrieve the product list"
// @Router /admin/inventories [get]
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

	message := "product list ended"

	successRes := response.ClientResponse(http.StatusOK, message, products_list, nil)
	c.JSON(http.StatusOK, successRes)
}


func (u *InventoryHandler) EditInventory(c *gin.Context) {
	var inventory domain.Inventories

	id := c.Query("inventory_id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "problems in the id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := c.BindJSON(&inventory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	modInventory, err := u.InventoryUseCase.EditInventory(inventory, idInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not edit the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sucessfully edited products", modInventory, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Delete Inventory
// @Description Delete an existing inventory item.
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Tags Admin Product Management
// @Param id query string true "ID of the inventory item to delete"
// @Success 200 {object} response.Response "Inventory item successfully deleted"
// @Failure 400 {object} response.Response "Invalid request format or fields provided in the wrong format"
// @Failure 404 {object} response.Response "Inventory item not found"
// @Failure 500 {object} response.Response "Failed to delete the inventory item"
// @Router /admin/inventories [delete]
func (u *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")

	err := u.InventoryUseCase.DeleteInventory(inventoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Sucessfully deleted the product", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Update Inventory
// @Description Update the stock of an existing inventory item.
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Tags Admin Product Management
// @Param body body models.InventoryUpdate true "Inventory update details in JSON format"
// @Success 200 {object} response.Response "Inventory stock successfully updated"
// @Failure 400 {object} response.Response "Invalid request format or fields provided in the wrong format"
// @Failure 404 {object} response.Response "Inventory item not found"
// @Failure 500 {object} response.Response "Failed to update the inventory stock"
// @Router /admin/inventories/stock [put]
func (i *InventoryHandler) UpdateInventory(c *gin.Context) {

	var p models.InventoryUpdate

	if err := c.BindJSON(&p); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fileds are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	a, err := i.InventoryUseCase.UpdateInventory(p.Productid, p.Stock)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could  not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Sucessfully upadated inventory stock", a, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) ShowIndividualProducts(c *gin.Context) {

	id := c.Query("id")
	product, err := i.InventoryUseCase.ShowIndividualProducts(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Search Products
// @Description Retrieve product details based on a prefix search for the product name.
// @Accept json
// @Produce json
// @Tags User Product Management
// @Param body body models.SearchItems true "Prefix for product name search"
// @Success 200 {object} response.Response "Product details retrieved successfully"
// @Failure 400 {object} response.Response "Invalid request format or fields provided in the wrong format"
// @Failure 500 {object} response.Response "Could not retrieve products by prefix search"
// @Router /user/product/search [post]
func (i *InventoryHandler) SearchProducts(c *gin.Context) {

	var prefix models.SearchItems

	if err := c.ShouldBindJSON(&prefix); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	productDetails, err := i.InventoryUseCase.SearchProductsOnPrefix(prefix.ProductName)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not retrive products by prefix search", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrived all details", productDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Filter Products by Category
// @Description Retrieve a list of products based on the specified category ID.
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Tags User Product Management
// @Param category_id query int true "Category ID for filtering products"
// @Success 200 {object} response.Response "Products list retrieved successfully"
// @Failure 400 {object} response.Response "Invalid category ID or products cannot be displayed"
// @Failure 500 {object} response.Response "Error in retrieving products"
// @Router /user/product/filter [get]
func (i *InventoryHandler) FilterCategory(c *gin.Context) {

	CategoryId := c.Query("category_id")
	CategoryIdInt, err := strconv.Atoi(CategoryId)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "products Cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	product_list, err := i.InventoryUseCase.FilterByCategory(CategoryIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "products cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	sucessRes := response.ClientResponse(http.StatusOK, "Products List", product_list, nil)
	c.JSON(http.StatusOK, sucessRes)

}
