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

type DashBoardUser struct {
	TotalUsers  int `json:"Totaluser"`
	BlockedUser int `json:"Blockuser"`
}
type DashBoardProduct struct {
	TotalProducts     int `json:"Totalproduct"`
	OutofStockProduct int `json:"Outofstock"`
}
type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type CompleteAdminDashboard struct {
	DashboardUser    DashBoardUser
	DashboardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashboardRevenue
	DashboardAmount  DashboardAmount
}
type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	CancelledOrder  int
	ReturnOrder     int
	TrendingProduct string
}
