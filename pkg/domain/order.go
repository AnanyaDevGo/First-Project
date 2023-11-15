package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint    `json:"user_id" gorm:"not null"`
	Users       Users   `json:"-" gorm:"foreignkey:UserID"`
	AddressID   uint    `json:"address_id" gorm:"not null"`
	Address     Address `json:"-" gorm:"foreignkey:AddressID"`
	FinalPrice  float64 `json:"price"`
	OrderStatus string  `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
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

type AdminOrdersResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
	Returned  []OrderDetails
}

type OrderDetails struct {
	ID       int     `json:"id" gorm:"id"`
	Username string  `json:"name"`
	Address  string  `json:"address"`
	Total    float64 `json:"total"`
}
