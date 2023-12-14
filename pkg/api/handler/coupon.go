package handler

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponUseCase interfaces.CouponUseCase
}

func NewCouponHandler(coupon interfaces.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		couponUseCase: coupon,
	}
}
func (cu *CouponHandler) AddCoupon(c *gin.Context) {
	var coupon domain.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	couponRes, err := cu.couponUseCase.AddCoupon(coupon)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot add coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added coupon", couponRes, nil)
	c.JSON(http.StatusOK, successRes)
}

func (cu *CouponHandler) GetCoupon(c *gin.Context) {
	couponRes, err := cu.couponUseCase.GetCoupon()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all coupons", couponRes, nil)
	c.JSON(http.StatusOK, successRes)
}

func (cu *CouponHandler) EditCoupon(c *gin.Context) {
	var edit domain.Coupon
	if err := c.BindJSON(&edit); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	fmt.Println("coupon ........", edit.IsAvailable)
	couponRes, err := cu.couponUseCase.EditCoupon(edit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot edit coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited coupon", couponRes, nil)
	c.JSON(http.StatusOK, successRes)
}
