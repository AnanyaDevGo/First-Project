package domain

type Coupon struct {
	Id                 uint   `json:"id" gorm:"primarykey"`
	Name               string `json:"name" gorm:"unique"`
	IsAvailable        bool   `json:"is_available" gorm:"default:false"`
	DiscountPercentage int    `json:"discount_percentage" gorm:"default:5"`
	MinimumPrice       int    `json:"minimum_price" gorm:"default:500"`
}
