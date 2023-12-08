package models

type CouponResp struct {
	Id                 uint   `json:"id"`
	Name               string `json:"name"`
	IsAvailable        bool   `json:"available"`
	DiscountPercentage int    `json:"discount_percentage"`
	MinimumPrice       int    `json:"minimum_price"`
}
