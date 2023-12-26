package handler

import (
	"CrocsClub/pkg/helper"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
	"strconv"
	"time"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary Admin Login
// @Description Authenticate an admin user and get access and refresh tokens.
// @Accept json
// @Produce json
// @Tags admin
// @Param body body models.AdminLogin true "Admin login details in JSON format"
// @Success 200 {object} response.Response "Admin authenticated successfully"
// @Failure 400 {object} response.Response  "Cannot authenticate user"
// @Router /admin/login [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	fmt.Println("it is here")
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Set("Access", admin.AccessToken)
	c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Admin Dashboard
// @Description Display the admin dashboard.
// @Accept json
// @Produce json
// @Tags admin
// @Success 200 {object} response.Response "Admin dashboard displayed"
// @Failure 500 {object} response.Response "Dashboard could not be displayed"
// @Router /admin/dashboard [get]
func (ad *AdminHandler) Dashboard(c *gin.Context) {

	dashboard, err := ad.adminUseCase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed", dashboard, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Filtered Sales Report
// @Description Retrieve sales report based on the specified time period.
// @Accept json
// @Produce json
// @Tags admin
// @Param period query string true "Time period for filtering sales report"
// @Success 200 {object} response.Response "Sales report retrieved successfully"
// @Failure 500 {object} response.Response "Sales report could not be retrieved"
// @Router /admin/sales-report [get]
func (ad *AdminHandler) FilteredSalesReport(c *gin.Context) {

	timePeriod := c.Query("period")
	salesReport, err := ad.adminUseCase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Block User
// @Description Block a user by their ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id query string true "User ID to be blocked"
// @Success 200 {object} response.Response "Successfully blocked the user"
// @Failure 400 {object} response.Response "User could not be blocked"
// @Router /admin/block-user [post]
func (ad *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Query("id")
	err := ad.adminUseCase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Unblock User
// @Description Unblock a user by their ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id query string true "User ID to be unblocked"
// @Success 200 {object} response.Response "Successfully unblocked the user"
// @Failure 400 {object} response.Response "User could not be unblocked"
// @Router /admin/unblock-user [post]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {

	id := c.Query("id")
	err := ad.adminUseCase.UnBlockUser(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Get Users
// @Description Retrieve a paginated list of users.
// @Accept json
// @Produce json
// @Tags admin
// @Param page query int true "Page number for pagination"
// @Success 200 {object} response.Response "Successfully retrieved the users"
// @Failure 400 {object} response.Response "Page number not in the right format or could not retrieve records"
// @Router /admin/get-users [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.adminUseCase.GetUsers(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary New Payment Method
// @Description Add a new payment method.
// @Accept json
// @Produce json
// @Tags admin
// @Param body body models.NewPaymentMethod true "Payment method details in JSON format"
// @Success 200 {object} response.Response "Successfully added Payment Method"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or could not add the payment method"
// @Router /admin/new-payment-method [post]
func (i *AdminHandler) NewPaymentMethod(c *gin.Context) {

	var method models.NewPaymentMethod
	if err := c.BindJSON(&method); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := i.adminUseCase.NewPaymentMethod(method.PaymentMethod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Payment Method", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary List Payment Methods
// @Description Retrieve a list of all available payment methods.
// @Accept json
// @Produce json
// @Tags admin
// @Success 200 {object} response.Response "Successfully got all payment methods"
// @Failure 500 {object} response.Response "Fields provided are in the wrong format or could not retrieve payment methods"
// @Router /admin/list-payment-methods [get]
func (a *AdminHandler) ListPaymentMethods(c *gin.Context) {

	categories, err := a.adminUseCase.ListPaymentMethods()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all payment methods", categories, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Delete Payment Method
// @Description Delete a payment method by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id query int true "Payment method ID to be deleted"
// @Success 200 {object} response.Response "Successfully deleted the payment method"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or error in deleting data"
// @Router /admin/delete-payment-method [delete]
func (a *AdminHandler) DeletePaymentMethod(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = a.adminUseCase.DeletePaymentMethod(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "error in deleting data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the Category", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Validate Refresh Token and Create New Access Token
// @Description Validate the provided refresh token and generate a new access token.
// @Accept json
// @Produce json
// @Tags admin
// @Security ApiKeyAuth
// @Param RefreshToken header string true "Refresh Token for validation" default("refresh_token_value")
// @Success 200 {string} string "New access token generated successfully"
// @Failure 401 {object} response.Response "Refresh token is invalid: user has to login again"
// @Failure 500 {object} response.Response "Error in creating new access token"
// @Router /admin/validate-refresh-token [post]
func (a *AdminHandler) ValidateRefreshTokenAndCreateNewAccess(c *gin.Context) {

	refreshToken := c.Request.Header.Get("RefreshToken")

	_, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("refreshsecret"), nil
	})
	if err != nil {
		c.AbortWithError(401, errors.New("refresh token is invalid:user have to login again"))
		return
	}

	claims := &helper.AuthCustomClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newAccessToken, err := token.SignedString([]byte("accesssecret"))
	if err != nil {
		c.AbortWithError(500, errors.New("error in creating new access token"))
	}

	c.JSON(200, newAccessToken)
}

// @Summary Sales Report by Date
// @Description Retrieve sales report data within a specified date range.
// @Accept json
// @Produce json
// @Tags admin
// @Param start query string true "Start date in the format DD-MM-YYYY" default("01-01-2023")
// @Param end query string true "End date in the format DD-MM-YYYY" default("31-12-2023")
// @Success 200 {object} response.Response "Sales report retrieved successfully"
// @Failure 400 {object} response.Response "Start or end date is empty, start date conversion failed, end date conversion failed, or invalid date range"
// @Failure 500 {object} response.Response "Sales report could not be retrieved"
// @Router /admin/sales-report-by-date [get]
func (ad *AdminHandler) SalesReportByDate(c *gin.Context) {
	startDateStr := c.Query("start")
	endDateStr := c.Query("end")
	if startDateStr == "" || endDateStr == "" {
		err := response.ClientResponse(http.StatusBadRequest, "start or end date is empty", nil, "Empty date string")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	startDate, err := time.Parse("02-01-2006", startDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "start date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	endDate, err := time.Parse("02-01-2006", endDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "end date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if startDate.After(endDate) {
		err := response.ClientResponse(http.StatusBadRequest, "start date is after end date", nil, "Invalid date range")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	report, err := ad.adminUseCase.ExecuteSalesReportByDate(startDate, endDate)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	success := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", report, nil)
	c.JSON(http.StatusOK, success)
}
