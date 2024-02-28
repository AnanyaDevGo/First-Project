package handler

import (
	"CrocsClub/pkg/helper"
	interfaces "CrocsClub/pkg/helper/interface"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
	"strconv"
	"time"


	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
	helper       interfaces.Helper
}

func NewAdminHandler(usecase services.AdminUseCase, helper interfaces.Helper) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
		helper:       helper,
	}
}

// @Summary Admin Login
// @Description Authenticate an admin user and get access and refresh tokens.
// @Accept json
// @Produce json
// @Tags Admin
// @Param body body models.AdminLogin true "Admin login details in JSON format"
// @Success 200 {object} response.Response "Admin authenticated successfully"
// @Failure 400 {object} response.Response  "Cannot authenticate user"
// @Router /admin/adminlogin [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
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

// Dashboard retrieves the admin dashboard information.
//
// @Summary Retrieve admin dashboard information
// @Description Retrieve information for the admin dashboard
// @Tags Admin Dashboard
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Admin dashboard information retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
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

// FilteredSalesReport retrieves a filtered sales report based on the specified time period.
//
// @Summary Retrieve filtered sales report
// @Description Retrieve a sales report filtered by the specified time period
// @Tags Admin Dashboard
// @security BearerTokenAuth
// @Param period query string true "Time period for filtering the sales report"
// @Success 200 {object} response.Response "Sales report retrieved successfully"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/salesreport [get]
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

// BlockUser blocks a user by ID.
//
// @Summary Block a user
// @Description Block a user by its ID
// @Tags Admin User Management
// @security BearerTokenAuth
// @Param id query string true "User ID to block"
// @Success 200 {object} response.Response "Successfully blocked the user"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/users/block [put]
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

// UnBlockUser unblocks a user by ID.
//
// @Summary Unblock a user
// @Description Unblock a user by its ID
// @Tags Admin User Management
// @security BearerTokenAuth
// @Param id query string true "User ID to unblock"
// @Success 200 {object} response.Response "Successfully unblocked the user"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/users/unblock [put]
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

// GetUsers retrieves a list of users with optional pagination.
//
// @Summary Retrieve a list of users
// @Description Retrieve a list of users with optional pagination.
// @Tags Admin User Management
// @security BearerTokenAuth
// @Param page query integer false "Page number for pagination (default: 1)"
// @Success 200 {object} response.Response "Successfully retrieved the list of users"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/users [get]
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

// NewPaymentMethod adds a new payment method.
//
// @Summary Add a new payment method
// @Description Add a new payment method using the provided details
// @Tags Admin Payment Management
// @security BearerTokenAuth
// @Accept json
// @Produce json
// @Param method body models.NewPaymentMethod true "New payment method details"
// @Success 200 {object} response.Response "Successfully added the payment method"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/payment-method/pay [post]
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

// ListPaymentMethods gets a list of payment methods.
// @Summary Get a list of payment methods
// @Description Get a list of all payment methods
// @Tags Admin Payment Management
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved the list of payment methods"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/payment-method [get]
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

// DeletePaymentMethod deletes a payment method by ID.
// @Summary Delete a payment method
// @Description Delete a payment method by its ID
// @Tags Admin Payment Management
// @security BearerTokenAuth
// @Param id query integer true "Payment method ID to delete"
// @Success 200 {object} response.Response "Successfully deleted the payment method"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/payment-method [delete]
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

// SalesReportByDate generates a sales report based on the specified start and end dates.
//
// @Summary Generate a sales report by date range
// @Description Generate a sales report based on the specified start and end dates.
// @Tags Admin Dashboard
// @security BearerTokenAuth
// @Param start query string true "Start date (format: DD-MM-YYYY)"
// @Param end query string true "End date (format: DD-MM-YYYY)"
// @Success 200 {object} response.Response "Successfully retrieved the sales report"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/sales-report-date [get]
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

// SalesByDate gets sales details for a specific date and allows downloading the report in PDF or Excel format.
//
// @Summary Get sales details by date
// @Description Get sales details for a specific date and download the report in PDF or Excel format
// @Tags Admin DashBoard
// @security BearerTokenAuth
// @Param year query integer true "Year for sales data"
// @Param month query integer true "Month for sales data"
// @Param day query integer true "Day for sales data"
// @Param download query string false "Download format (pdf or excel)"
// @Success 200 {object} response.Response "Successfully retrieved sales details"
// @Failure 400 {object} response.Response "Invalid request or incorrect format"
// @Failure 502 {object} response.Response "Bad Gateway"
// @Router /admin/salesbydate [get]
func (a *AdminHandler) SalesByDate(c *gin.Context) {
	year := c.Query("year")
	yearInt, err := strconv.Atoi(year)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting year", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	month := c.Query("month")
	monthInt, err := strconv.Atoi(month)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting month", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	day := c.Query("day")
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting day", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := a.adminUseCase.SalesByDate(dayInt, monthInt, yearInt)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting sales details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	download := c.Query("download")
	if download == "pdf" {
		pdf, err := a.adminUseCase.PrintSalesReport(body)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		c.Header("Content-Disposition", "attachment;filename=totalsalesreport.pdf")

		pdfFilePath := "salesReport/totalsalesreport.pdf"

		err = pdf.OutputFileAndClose(pdfFilePath)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		c.Header("Content-Disposition", "attachment; filename=total_sales_report.pdf")
		c.Header("Content-Type", "application/pdf")

		c.File(pdfFilePath)

		c.Header("Content-Type", "application/pdf")

		err = pdf.Output(c.Writer)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
	} else {
		excel, err := a.helper.ConvertToExel(body)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		fileName := "sales_report.xlsx"

		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		if err := excel.Write(c.Writer); err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "Error in serving the sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
	}

	succesRes := response.ClientResponse(http.StatusOK, "success", body, nil)
	c.JSON(http.StatusOK, succesRes)
}
