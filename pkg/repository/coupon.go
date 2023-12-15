package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"

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

func (cr *couponRepository) EditCoupon(Edit domain.Coupon) (models.CouponResp, error) {
	query := "UPDATE coupons SET name=?, is_available=?, discount_percentage=?, minimum_price=? WHERE id=?"
	result := cr.DB.Exec(query, Edit.Name, Edit.IsAvailable, Edit.DiscountPercentage, Edit.MinimumPrice, Edit.Id)

	// Check for errors during the update
	if result.Error != nil {
		return models.CouponResp{}, fmt.Errorf("error in updating data: %v", result.Error)
	}

	// Check if any rows were affected by the update
	if result.RowsAffected == 0 {
		return models.CouponResp{}, errors.New("no rows were affected, coupon with the specified ID may not exist")
	}

	// Create a models.CouponResp instance from domain.Coupon
	couponResp := models.CouponResp{
		Id:                 Edit.Id,
		Name:               Edit.Name,
		IsAvailable:        Edit.IsAvailable,
		DiscountPercentage: Edit.DiscountPercentage,
		MinimumPrice:       Edit.MinimumPrice,
	}

	return couponResp, nil
}

func (cr *couponRepository) CheckCouponById(CouponId int) (bool, error) {
	var count int
	if err := cr.DB.Raw("select count(*) from coupons where id=?", CouponId).Scan(&count).Error; err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}
func (cr *couponRepository) GetCouponById(CouponId int) (models.CouponResp, error) {
	var coupon models.CouponResp
	if err := cr.DB.Raw("select * from coupons where id=?", CouponId).Scan(&coupon).Error; err != nil {
		return models.CouponResp{}, errors.New("cannot retrive data")
	}
	return coupon, nil
}
