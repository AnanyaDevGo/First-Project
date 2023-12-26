package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	usecase services.PaymentUseCase
}

func NewPaymentHandler(use services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: use,
	}
}

// @Summary Make Payment using RazorPay
// @Description Generate order details and initiate payment using RazorPay.
// @Accept json
// @Produce json
// @Tags payment
// @Param id query string true "Order ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} response.ResponseHTML "HTML page with RazorPay payment details"
// @Failure 400 {object} response.Response "Invalid or missing parameters"
// @Failure 500 {object} response.Response "Failed to generate order details"
// @Router /payment/razorpay [get]
func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {

	orderID := c.Query("id")
	userID := c.Query("user_id")
	usrid, _ := strconv.Atoi(userID)

	orderDetail, err := p.usecase.MakePaymentRazorPay(orderID, usrid)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

// @Summary Verify Payment
// @Description Verify and update payment details after completing the RazorPay payment.
// @Accept json
// @Produce json
// @Tags payment
// @Param order_id query string true "Order ID"
// @Param payment_id query string true "Payment ID"
// @Param razor_id query string true "RazorPay ID"
// @Success 200 {object} response.Response "Successfully updated payment details"
// @Failure 400 {object} response.Response "Invalid or missing parameters"
// @Failure 500 {object} response.Response "Failed to update payment details"
// @Router /payment/verify [get]
func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	err := p.usecase.VerifyPayment(paymentID, razorID, orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
