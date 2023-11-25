package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(userid int, addressid int, paymentid int) error
	GetOrders(orderId int) (domain.OrderResponse, error)
	CancelOrder(orderId int) error

	GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error)
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
}
