package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
	IsDeleted    bool   `json:"is_deleted" gorm:"default:false"`
}

type Order struct {
	gorm.Model
	UserID          uint          `json:"user_id" gorm:"not null"`
	Users           Users         `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"price"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:5;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED','Return To Wallet')"`
	PaymentStatus   string        `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID'"`
	Approval        bool          `json:"approval" gorm:"default:false"`
}

type OrderResponse struct {
	UserID          uint          `json:"user_id" gorm:"not null"`
	Users           Users         `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"price"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string        `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID')"`
}

type OrderItem struct {
	ID          uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint        `json:"order_id"`
	Order       Order       `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	InventoryID uint        `json:"inventory_id"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int         `json:"quantity"`
	TotalPrice  float64     `json:"total_price"`
}

type CombinedOrderDetails struct {
	OrderId       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	OrderStatus   string  `json:"order_status"`
	PaymentStatus bool    `json:"payment_status"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
	Street        string  `json:"street"`
	City          string  `json:"city"`
}

type OrderDetails struct {
	ID            int     `json:"id" gorm:"id"`
	Username      string  `json:"name"`
	Address       string  `json:"address"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	Total         float64 `json:"total"`
}

type OrderDetailsWithImages struct {
	OrderDetails  Order
	Images        []string
	PaymentMethod string
}

type AdminOrdersResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
	Returned  []OrderDetails
}
