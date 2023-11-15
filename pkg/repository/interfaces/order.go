package interfaces

import "CrocsClub/pkg/domain"

type OrderRepository interface {
	GetOrders(id int) ([]domain.Order, error)
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	CheckOrderStatusByID(id int) (string, error)
}
