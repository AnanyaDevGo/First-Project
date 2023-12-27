package handler

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
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

// @Summary Add Coupon
// @Description Add a new coupon.
// @Accept json
// @Produce json
// @Tags Admin Coupon Management
// @security BearerTokenAuth
// @Param body body domain.Coupon true "Coupon details in JSON format"
// @Success 200 {object} response.Response "Successfully added coupon"
// @Failure 400 {object} response.Response "Field provided in the wrong format or Cannot add coupon"
// @Router /admin/coupon [post]
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

// @Summary Get Coupons
// @Description Retrieve all coupons.
// @Accept json
// @Produce json
// @Tags Admin Coupon Management
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully got all coupons"
// @Failure 400 {object} response.Response "Error in getting coupons"
// @Router /admin/coupon [get]
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

// @Summary Edit Coupon
// @Description Edit an existing coupon.
// @Accept json
// @Produce json
// @Tags Admin Coupon Management
// @security BearerTokenAuth
// @Param body body domain.Coupon true "Coupon details in JSON format"
// @Success 200 {object} response.Response "Successfully edited coupon"
// @Failure 400 {object} response.Response "Field provided in the wrong format or Cannot edit coupon"
// @Router /admin/coupon [patch]
func (cu *CouponHandler) EditCoupon(c *gin.Context) {
	var edit domain.Coupon
	if err := c.BindJSON(&edit); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	couponRes, err := cu.couponUseCase.EditCoupon(edit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot edit coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited coupon", couponRes, nil)
	c.JSON(http.StatusOK, successRes)
}
