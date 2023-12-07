package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"

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
func (ch *CartHandler) AddToCart(c *gin.Context) {
	// Get user ID from the context
	idString, exists := c.Get("id")
	if !exists {
		err := errors.New("user ID not found in the context")
		errorRes := response.ClientResponse(http.StatusBadRequest, "User ID not found in the context", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println("userid at handler", idString)
	userID, ok := idString.(int)
	if !ok {
		errorRes := response.ClientResponse(http.StatusBadRequest, "User ID not in the right format", nil, "")
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var cart models.Cart
	if err := c.BindJSON(&cart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to parse request JSON", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	fmt.Println("inventory id ", cart.InventoryID)
	fmt.Println("quantity", cart.Quantity)

	if err := ch.usecase.AddToCart(userID, cart.InventoryID, cart.Quantity); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add to the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *CartHandler) CheckOut(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	products, err := i.usecase.CheckOut(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
