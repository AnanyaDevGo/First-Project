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
