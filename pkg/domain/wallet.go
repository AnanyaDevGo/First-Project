package domain

type Wallet struct {
	ID     uint    `json:"id"  gorm:"unique;not null"`
	UserID int     `json:"user_id"`
	Users  Users   `json:"-" gorm:"foreignkey:UserID"`
	Amount float64 `json:"amount" gorm:"default:0"`
}
type WalletHistory struct {
	ID         uint    `json:"id"  gorm:"unique;not null"`
	WalletID   int     `json:"wallet_id"`
	Wallet     Wallet  `json:"-" gorm:"foreignkey:WalletID"`
	OrderID    int     `json:"order_id"`
	Amount     float64 `json:"amount"`
	IsCredited bool    `json:"is_credited" gorm:"default:true"`
}
