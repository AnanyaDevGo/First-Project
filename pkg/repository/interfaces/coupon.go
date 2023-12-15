package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(coupon domain.Coupon) (models.CouponResp, error)
	CouponExist(name string) (bool, error)
	GetCoupon() ([]models.CouponResp, error)
	EditCoupon(Edit domain.Coupon) (models.CouponResp, error)
	CheckCouponById(CouponId int) (bool, error)
	GetCouponById(CouponId int) (models.CouponResp, error)
}
