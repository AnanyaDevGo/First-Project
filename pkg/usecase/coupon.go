package usecase

import (
	"CrocsClub/pkg/domain"
	repo "CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
)

type couponUseCase struct {
	couponRepo repo.CouponRepository
}

func NewCouponUsecase(couponRepo repo.CouponRepository) interfaces.CouponUseCase {
	return &couponUseCase{
		couponRepo: couponRepo,
	}
}

func (cu *couponUseCase) AddCoupon(coupon domain.Coupon) (models.CouponResp, error) {
	if coupon.Name == "" {
		return models.CouponResp{}, errors.New("cannot put empty values in name")
	}
	if coupon.DiscountPercentage <= 0 || coupon.MinimumPrice <= 0 {
		return models.CouponResp{}, errors.New("values cannot be negative")
	}
	ok, err := cu.couponRepo.CouponExist(coupon.Name)
	if err != nil {
		return models.CouponResp{}, errors.New("failed to get coupon details")
	}
	if ok {
		return models.CouponResp{}, errors.New("coupon already exist")
	}
	couponResp, err := cu.couponRepo.AddCoupon(coupon)

	if err != nil {
		return models.CouponResp{}, errors.New("adding coupon failed")
	}
	return couponResp, nil
}
func (cu *couponUseCase) GetCoupon() ([]models.CouponResp, error) {
	couponResp, err := cu.couponRepo.GetCoupon()
	if err != nil {
		return []models.CouponResp{}, err
	}
	return couponResp, nil
}
func (cu *couponUseCase) EditCoupon(Edit models.CouponResp) (models.CouponResp, error){

	if Edit.Name == "" {
		return models.CouponResp{}, errors.New("cannot put empty values in name")
	}
	if Edit.DiscountPercentage <1 || Edit.MinimumPrice <1 {
		return models.CouponResp{}, errors.New("cannot put negative values")
	}
	ok, err := cu.couponRepo.CheckCouponById(int(Edit.Id))
	if err != nil {
		return models.CouponResp{},errors.New("failed to get coupon details")
	}
	if ok{
		return models.CouponResp{},errors.New("coupon already exist")
	}
	couponResp, err := cu.couponRepo.EditCoupon(Edit)
	if err != nil {
		return models.CouponResp{}, errors.New("failed to edit coupon")
	}
	return couponResp, nil
}