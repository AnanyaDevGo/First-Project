package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
	"time"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)

	NewPaymentMethod(string) error
	GetPaymentMethod() ([]models.PaymentMethodResponse, error)
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	CheckIfPaymentMethodIdExists(payment string) (bool, error)
	CheckIfPaymentMethodNameExists(payment string) (bool, error)
	DeletePaymentMethod(id int) error

	TotalRevenue() (models.DashboardRevenue, error)
	DashBoardOrder() (models.DashboardOrder, error)
	AmountDetails() (models.DashboardAmount, error)
	DashBoardUserDetails() (models.DashBoardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)
	FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
	SalesByDay(yearInt int, monthInt int, dayInt int) ([]models.OrderDetailsAdmin, error)
	SalesByMonth(yearInt int, monthInt int) ([]models.OrderDetailsAdmin, error)
	SalesByYear(yearInt int) ([]models.OrderDetailsAdmin, error)
}
