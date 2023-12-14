package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
	"fmt"
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

func (u *UserHandler) GetCart(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	fmt.Println("card id", id)

	products, err := u.userUseCase.GetCart(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

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

func (i *UserHandler) UpdateQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("here")
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properlyyyyy", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	inv, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		fmt.Println("****here****")
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
