package models

type ProductsReceiver struct {
	CategoryID  uint    `json:"Category_id"`
	ProductName string  `json:"product_name"`
	Size_id     uint    `json:"size_id"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category_name"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
type ProductsResponse struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  string   `json:"category_id"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
}
type ProductsResponseDisp struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  string   `json:"category_id"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
}
type ProductsResp struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  string   `json:"category_id"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
}
type MakeOrder struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
}
type Order struct {
	UserID          int  `json:"user_id"`
	AddressID       int  `json:"address_id"`
	PaymentMethodID int  `json:"payment_id"`
	CouponID        int  `json:"coupon_id"`
	UseWallet       bool `json:"use_wallet" gorm:"default:false"`
}
type SearchItems struct {
	ProductName string `json:"product_name"`
}
type ItemOrderDetails struct {
	OrderId       string  `json:"order_id"`
	FinalPrice    float64 `json:"final_price"`
	OrderStatus   string  `json:"order_status" gorm:"column:order_status"`
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
type ItemDetails struct {
	ProductName string  `json:"product_name"`
	FinalPrice  float64 `json:"final_price"`
	Price       float64 `json:"price" `
	Total       float64 `json:"total_price"`
	Quantity    int     `json:"quantity"`
}
type CatRes struct {
	Category string `json:"category_name"`
}
