package models

type OrderDetails struct {
	ID        int    `json:"order_id"`
	UserName  string `json:"name"`
	AddressID int    `json:"address_id"`
}

type CombinedOrderDetails struct {
	OrderId        string  `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
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
	CouponUsed  string
	OrderStatus string
}

type ProductDetails struct {
	ProductName string
	Image       string
	Quantity    int
	Amount      float64
}
