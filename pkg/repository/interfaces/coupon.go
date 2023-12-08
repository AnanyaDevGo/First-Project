package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(coupon domain.Coupon) (models.CouponResp, error)
	CouponExist(name string)(bool, error)
}
