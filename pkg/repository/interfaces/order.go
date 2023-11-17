package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OrderRepository interface {
	GetOrders(id int) ([]domain.Order, error)
	CancelOrder(id int) error
	OrderItems(userid, addressid, paymentid int, total float64, coupon string) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	EditOrderStatus(status string, id int) error
	CheckOrderStatusByID(id int) (string, error)
	AdminOrders(status string) ([]domain.OrderDetails, error)
	GetOrderDetail(orderID string) (domain.Order, error)
	MakePaymentStatusAsPaid(id int) error
}
