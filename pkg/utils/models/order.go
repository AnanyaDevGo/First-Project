package models

type OrderDetails struct {
	ID int `json:"id"`
	AddressID       int     `json:"address_id"`
	PaymentMethodID int     `json:"payment_method_id"`
	FinalPrice      float64 `json:"final_price"`
	OrderStatus     string  `json:"order_status" gorm:"column:order_status"`
}

type CombinedOrderDetails struct {
	OrderId       string  `json:"order_id"`
	FinalPrice    float64 `json:"final_price"`
	PaymentStatus string  `json:"payment_status"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
	Street        string  `json:"street"`
	City          string  `json:"city"`
}

type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `josn:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type EditOrderStatus struct {
	OrderID int    `json:"order_id"`
	Status  string `json:"order_status"`
}

type IndividualOrderDetails struct {
	OrderID     int
	Address     string
	Phone       string
	Products    []ProductDetails `gorm:"-"`
	TotalAmount float64
	OrderStatus string
}

type ProductDetails struct {
	ProductName string
	Image       string
	Quantity    int
	Amount      float64
}

type PaymentMethodResponse struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}
type ReturnOrderResponse struct {
	UserId      int    `json:"user_id"`
	OrderID     int    `json:"order_id"`
	OrderStatus string `json:"order_status"`
	CartAmount  int    `jsom:"cart_amount"`
}

type OrderDetailsRep struct {
	ID int `json:"order_id"`
	// UserName      string  `json:"name"`
	// AddressID     int     `json:"address_id"`
	PaymentMethod int     `json:"payment_method_id"`
	Total         int     `json:"total_amount"`
	FinalPrice    float64 `json:"final_price"`
	// OrderStatus   string  `json:"order_status" gorm:"column:order_status"`
}
type OrderDetailsAdmin struct {
	TotalAmount float64 `gorm:"column:total_amount"`
	ProductName string  `gorm:"column:product_name"`
}
