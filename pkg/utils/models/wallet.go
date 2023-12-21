package models

type WalletAmount struct {
	Amount float64 `json:"amount"`
}

type WalletHistory struct {
	ID         int     `json:"id"  gorm:"unique;not null"`
	OrderID    int     `json:"order_id"`
	Amount     float64 `json:"amount"`
	IsCredited bool    `json:"is_credited"`
}
