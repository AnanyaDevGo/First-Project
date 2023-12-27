package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OrderRepository interface {
	GetOrders(orderId int) (domain.OrderResponse, error)
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CheckOrderStatusByID(id int) (string, error)
	CancelOrder(orderId, userId, cartAmt int, paymentStatus string) error
	GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error)
	GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error)
	CheckOrdersStatusByID(id string) (string, error)
	GetShipmentStatus(orderId string) (string, error)
	ApproveOrder(orderId string) error
	ChangeOrderStatus(orderID, status string) error
	GetShipmentsStatus(orderID string) (string, error)
	PaymentMethodID(orderID int) (int, error)
	ReturnOrder(returnOrderResp models.ReturnOrderResponse) error
	ReduceInventoryQuantity(productName string, quantity int) error
	PaymentAlreadyPaid(orderID int) (bool, error)
	GetOrderDetailsByOrderId(orderID int) (models.CombinedOrderDetails, error)
	GetItemsByOrderId(orderId int) ([]models.ItemDetails, error)
	GetDetailedOrderThroughId(orderId int) (models.ItemOrderDetails, error)
	DebitWallet(userID int, Amount float64) error
}
