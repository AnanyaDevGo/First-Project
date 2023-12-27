package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
)

type OrderUseCase interface {
	OrderItemsFromCart(userid, addressid, paymentid, couponId int, useWallet bool) (models.OrderDetailsRep, error)
	GetOrders(orderId int) (domain.OrderResponse, error)
	GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error)
	CancelOrder(orderId int) error
	GetAdminOrders(page int) ([]models.CombinedOrderDetails, error)
	OrdersStatus(orderId string) error
	ReturnOrder(orderID string) error
	PaymentMethodID(order_id int) (int, error)
	PrintInvoice(orderIdInt int) (*gofpdf.Fpdf, error)
}
