package interfaces

import "CrocsClub/pkg/domain"

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
}
