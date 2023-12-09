package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type CouponUseCase interface {
	AddCoupon(coupon domain.Coupon) (models.CouponResp, error)
	GetCoupon() ([]models.CouponResp, error)
	EditCoupon(Edit models.CouponResp) (models.CouponResp, error)
}
