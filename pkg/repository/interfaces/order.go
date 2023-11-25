package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OrderRepository interface {
	GetOrders(orderId int) (domain.OrderResponse, error)
	CancelOrder(id int) error
	GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error)
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	EditOrderStatus(status string, id int) error
	CheckOrderStatusByID(id int) (string, error)
	AdminOrders(status string) ([]domain.OrderDetails, error)
	GetOrderDetail(orderID string) (domain.Order, error)
	GetProductImagesInAOrder(id int) ([]string, error)

	MakePaymentStatusAsPaid(id int) error
	FindPaymentMethodOfOrder(id int) (string, error)
}
