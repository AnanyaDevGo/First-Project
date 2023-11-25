package models

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}

type NewPaymentMethod struct {
	PaymentMethod string `json:"payment_method"`
}

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}

type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}

type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}

type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}

type DashboardUser struct {
	TotalUsers   int
	BlockedUser  int
	OrderedUsers int
}

type DashBoardProduct struct {
	TotalProducts     int
	OutOfStockProduct int
	TopSellingProduct string
}

type CompleteAdminDashboard struct {
	DashboardRevenue DashboardRevenue
	DashboardOrder   DashboardOrder
	DashboardAmount  DashboardAmount
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
}
