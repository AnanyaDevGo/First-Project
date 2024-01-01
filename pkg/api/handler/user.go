package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"

	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// UserSignUp handles user sign-up functionality.
// @Summary Register a new user
// @Description Create a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.UserDetails true "User details in JSON format"
// @Success 201 {object} models.UserDetails "User signed up successfully"
// @Failure 400 {object} models.TokenUsers "Invalid request or constraints not satisfied"
// @Failure 500 {object}  models.TokenUsers "Internal server error"
// @Router /user/signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {
	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest,
			errRes)
		return
	}
	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userCreated, err := u.userUseCase.UserSignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)
}

// LoginHandler handles the user login functionality.
// @Summary Log in a user
// @Description Logs in a user with provided credentials
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.UserLogin true "User login details in JSON format"
// @Success 200 {object} models.UserDetails "User details logged in successfully"
// @Failure 400 {object} response.Response "Invalid request or constraints not satisfied"
// @Failure 401 {object} models.UserDetails "Unauthorized: Invalid credentials"
// @Router /user/login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err := u.userUseCase.LoginHandler(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", user_details, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetCart retrieves the user's shopping cart.
//
// @Summary Retrieve user's shopping cart
// @Description Retrieve the products in the user's shopping cart
// @Tags User Cart Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Products in the shopping cart retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve the shopping cart"
// @Router /user/cart [get]
func (u *UserHandler) GetCart(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	products, err := u.userUseCase.GetCart(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// RemoveFromCart removes a product from the user's shopping cart.
//
// @Summary Remove product from cart
// @Description Remove a product from the user's shopping cart by specifying cart and inventory IDs
// @Tags User Cart Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Param inventory_id query integer true "ID of the inventory to be removed from the cart"
// @Success 200 {object} response.Response "Product successfully removed from the shopping cart"
// @Failure 400 {object} response.Response "Invalid request or missing parameters"
// @Failure 404 {object} response.Response "Product or cart not found"
// @Router /user/cart/remove [delete]
func (i *UserHandler) RemoveFromCart(c *gin.Context) {
	cartIDStr := c.Query("cart_id")
	if cartIDStr == "" {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cart_id is required", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "invalid cart_id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inventoryIDStr := c.Query("inventory_id")
	if inventoryIDStr == "" {
		errorRes := response.ClientResponse(http.StatusBadRequest, "inventory_id is required", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inventoryID, err := strconv.Atoi(inventoryIDStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "invalid inventory_id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.RemoveFromCart(cartID, inventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not remove from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Removed product from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// AddAddress adds a new address to the user's profile.
//
// @Summary Add new address
// @Description Add a new address to the user's profile by providing address details
// @Tags User Profile Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Param address body models.AddAddress true "Address details to be added"
// @Success 200 {object} response.Response "Address successfully added"
// @Failure 400 {object} response.Response "Invalid request or missing parameters"
// @Failure 403 {object} response.Response "Forbidden, user not authenticated"
// @Router /user/profile/addaddress [post]
func (u *UserHandler) AddAddress(c *gin.Context) {
	id, _ := c.Get("id")
	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := u.userUseCase.AddAddress(id.(int), address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdateQuantity updates the quantity of a product in the user's cart.
//
// @Summary Update product quantity
// @Description Update the quantity of a product in the user's cart
// @Tags User Cart Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Param inventory query integer true "Inventory ID of the product"
// @Param quantity query integer true "New quantity of the product"
// @Success 200 {object} response.Response "Quantity updated successfully"
// @Failure 400 {object} response.Response "Invalid request or missing parameters"
// @Failure 403 {object} response.Response "Forbidden, user not authenticated"
// @Router /user/cart [put]
func (i *UserHandler) UpdateQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properlyyyyy", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inv, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	qty, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.UpdateQuantity(id, inv, qty); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not Add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetAddress handles the retrieval of user addresses.
// @Summary Get user addresses
// @Description Retrieve user addresses by ID
// @Tags User Profile Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Success 200 {object} response.Response "User addresses retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve user addresses"
// @Router /user/profile/address [get]
func (u *UserHandler) GetAddress(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	addresses, err := u.userUseCase.GetAddress(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetUserDetails handles the retrieval of user details.
// @Summary Get user details
// @Description Retrieve user details by ID
// @Tags User Profile Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Success 200 {object} models.UserDetailsResponse "User details retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve user details"
// @Router /user/profile/details [get]
func (u *UserHandler) GetUserDetails(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	details, err := u.userUseCase.GetUserDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrive details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrived details", details, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Edit User Details
// @Description Edit details of the authenticated user.
// @Accept json
// @Produce json
// @Tags User Profile Management
// @security BearerTokenAuth
// @Param edit body models.Edit true "User details to be edited"
// @Success 201 {object} response.Response "Details edited successfully"
// @Failure 400 {object} response.Response "Invalid input or error updating values"
// @Router /user/profile/edit/ [patch]
func (u *UserHandler) Edit(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var model models.Edit
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(model)
	if err != nil {
		err = errors.New("missing constraints for email id")
		errRes := response.ClientResponse(http.StatusBadRequest, "email id is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	result, err := u.userUseCase.Edit(id, model)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error updating the values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "details edited succesfully", result, nil)

	c.JSON(http.StatusCreated, successRes)
}

// ChangePassword handles the changing of user password.
// @Summary Change user password
// @Description Change user password by ID
// @Tags User Profile Management
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Param changePasswordBody body models.ChangePassword true "Change password payload"
// @Success 200 {object} response.Response "Password changed successfully"
// @Failure 400 {object} response.Response "Failed to change user password"
// @Router /user/profile/security/change-password [put]
func (u *UserHandler) ChangePassword(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var ChangePassword models.ChangePassword
	if err := c.BindJSON(&ChangePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := u.userUseCase.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed Successfully ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
