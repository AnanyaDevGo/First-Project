package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase services.CartUseCase
}

func NewCartHandler(usecase services.CartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

func (i *CartHandler) AddToCart(c *gin.Context) {

	var model models.AddToCart
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.usecase.AddToCart(model.UserID, model.InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To cart", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *CartHandler) CheckOut(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.usecase.CheckOut(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
