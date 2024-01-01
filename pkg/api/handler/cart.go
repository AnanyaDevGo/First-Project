package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
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

// @Summary Add to Cart
// @Description Add a product to the user's shopping cart.
// @Accept json
// @Produce json
// @Tags User Cart Management
// @security BearerTokenAuth
// @Param cart body models.Cart true "Product details to be added to the cart in JSON format"
// @Success 200 {object} response.Response "Successfully added to cart"
// @Failure 400 {object} response.Response "User ID not found in the context, User ID not in the right format, Failed to parse request JSON, or Could not add to the cart"
// @Router /user/home/addcart [post]
func (ch *CartHandler) AddToCart(c *gin.Context) {
	// Get user ID from the context
	idString, exists := c.Get("id")
	if !exists {
		err := errors.New("user ID not found in the context")
		errorRes := response.ClientResponse(http.StatusBadRequest, "User ID not found in the context", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
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

	if err := ch.usecase.AddToCart(userID, cart.InventoryID, cart.Quantity); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add to the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Checkout
// @Description Process the checkout for the user and retrieve the checkout details.
// @Accept json
// @Produce json
// @Tags User Cart Management
// @security BearerTokenAuth
// @Param id header int true "User ID obtained from authentication"
// @Success 200 {object} response.Response "Successfully retrieved checkout details"
// @Failure 400 {object} response.Response "Could not open checkout"
// @Router /user/check-out [get]
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
