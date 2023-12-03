package domain

type Wallet struct {
	ID     int     `json:"id"  gorm:"unique;not null"`
	UserID int     `json:"user_id"`
	Users  Users   `json:"-" gorm:"foreignkey:UserID"`
	Amount float64 `json:"amount" gorm:"default:0"`
}
type WalletHistory struct {
	ID          int     `json:"id"  gorm:"unique;not null"`
	UserID      int     `json:"user_id"`
	OrderID     int     `json:"order_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	IsCredited  bool    `json:"is_credited" gorm:"default:true"`
}
