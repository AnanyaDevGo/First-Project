package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		DB: db,
	}
}

func (cr *couponRepository) AddCoupon(coupon domain.Coupon) (models.CouponResp, error) {

	var couponRep models.CouponResp

	query := `INSERT INTO coupons (name, is_available, discount_percentage, minimum_price) VALUES (?,?,?,?) RETURNING id,name,is_available, discount_percentage, minimum_price`

	err := cr.DB.Raw(query, coupon.Name, coupon.IsAvailable, coupon.DiscountPercentage, coupon.MinimumPrice).Scan(&couponRep).Error
	if err != nil {
		return models.CouponResp{}, errors.New("error in inserting")
	}
	return couponRep, nil
}

func (cr *couponRepository) CouponExist(name string) (bool, error) {
	var count int

	if err := cr.DB.Raw("select count(*) from coupons where name=?", name).Scan(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (cr *couponRepository) GetCoupon() ([]models.CouponResp, error) {
	var coupon []models.CouponResp
	if err := cr.DB.Raw("select * from coupons").Scan(&coupon).Error; err != nil {
		return []models.CouponResp{}, errors.New("cannot retrive data")
	}
	return coupon, nil
}
