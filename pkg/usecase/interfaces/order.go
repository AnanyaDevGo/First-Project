package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(userid, addressid, paymentid, couponId int) error
	GetOrders(orderId int) (domain.OrderResponse, error)
	GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error)
	CancelOrder(orderId int) error
	GetAdminOrders(page int) ([]models.CombinedOrderDetails, error)
	OrdersStatus(orderId string) error
	ReturnOrder(orderID string) error
	PaymentMethodID(order_id int) (int, error)
}
