package usecase

import (
	"CrocsClub/pkg/domain"
	repo "CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"
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
func (cu *couponUseCase) EditCoupon(edit domain.Coupon) (models.CouponResp, error) {
	if edit.Name == "" {
		return models.CouponResp{}, errors.New("name cannot be empty")
	}

	if edit.DiscountPercentage < 1 || edit.MinimumPrice < 1 {
		return models.CouponResp{}, errors.New("discount percentage and minimum price must be greater than or equal to 1")
	}

	fmt.Println("coupon avilable", edit.IsAvailable)
	// exists, err := cu.couponRepo.CheckCouponById(int(edit.Id))
	// if err != nil {
	// 	return models.CouponResp{}, fmt.Errorf("failed to check coupon details: %v", err)
	// }
	// if exists {
	// 	return models.CouponResp{}, errors.New("coupon with the specified ID already exists")
	// }

	couponResp, err := cu.couponRepo.EditCoupon(edit)
	if err != nil {
		return models.CouponResp{}, fmt.Errorf("failed to edit coupon: %v", err)
	}

	return couponResp, nil
}
// func (cu *couponUseCase) Applycoupon(coupon string, userId int)error{
// 	ok, err := cu.couponRepo.CouponExist(coupon)
// 	if err != nil{
// 		return err
// 	}
// 	if !ok {
// 		return errors.New("coupon does not exist")
// 	}
	
// }