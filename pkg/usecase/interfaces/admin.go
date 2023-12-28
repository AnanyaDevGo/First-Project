package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	NewPaymentMethod(string) error
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	DeletePaymentMethod(id int) error
	DashBoard() (models.CompleteAdminDashboard, error)
	FilteredSalesReport(timePeriod string) (models.SalesReport, error)
	ExecuteSalesReportByDate(startDate, endDate time.Time) (models.SalesReport, error)
	SalesByDate(dayInt int, monthInt int, yearInt int) ([]models.OrderDetailsAdmin, error)
	PrintSalesReport(sales []models.OrderDetailsAdmin) (*gofpdf.Fpdf, error)
}
