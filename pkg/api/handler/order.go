package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	id, ok := c.Get("id")
	if !ok {
		err := errors.New("error in getting userId")
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	userId, ok := id.(int)
	if !ok {
		err := errors.New("invalid id ")
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	var order models.Order
	order.UserID = userId
	order.CouponID = 0
	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID, order.CouponID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully made the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) GetOrders(c *gin.Context) {

	idString := c.Query("order_id")
	order_id, err := strconv.Atoi(idString)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetOrders(order_id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i OrderHandler) CancelOrder(c *gin.Context) {
	idString := c.Query("order_id")
	orderID, err := strconv.Atoi(idString)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = i.orderUseCase.CancelOrder(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not cancel the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order successfully canceled", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) GetAllOrders(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("count", "10"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	id, _ := c.Get("id")
	UserID, _ := id.(int)

	orders, err := i.orderUseCase.GetAllOrders(UserID, page, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) GetAdminOrders(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetAdminOrders(page)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) ApproveOrder(c *gin.Context) {
	orderID := c.Query("order_id")

	err := i.orderUseCase.OrdersStatus(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not approve order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully approved order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (o *OrderHandler) ReturnOrder(c *gin.Context) {

	orderID := c.Query("order_id")

	err := o.orderUseCase.ReturnOrder(orderID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "order could not be returned", nil, err)
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully returned", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (O *OrderHandler) PrintInvoice(c *gin.Context) {
	orderId := c.Query("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		err = errors.New("error in coverting order id" + err.Error())
		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the order id", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	pdf, err := O.orderUseCase.PrintInvoice(orderIdInt)
	fmt.Println("error ", err)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing the invoice", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.Header("Content-Disposition", "attachment;filename=invoice.pdf")

	pdfFilePath := "salesReport/invoice.pdf"

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	c.File(pdfFilePath)

	c.Header("Content-Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the request was succesful", pdf, nil)
	c.JSON(http.StatusOK, successRes)
}
