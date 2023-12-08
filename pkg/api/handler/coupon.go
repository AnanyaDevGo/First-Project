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
	fmt.Println("here...............0")
	var coupon domain.Coupon
	fmt.Println("here...............1")
	if err := c.BindJSON(&coupon); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		fmt.Println("here...............2")
		return
	}
	couponRes, err := cu.couponUseCase.AddCoupon(coupon)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot add coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		fmt.Println("here...............3")
		return
	}
	fmt.Println(".....................4", couponRes)
	successRes := response.ClientResponse(http.StatusOK, "successfully added coupon", couponRes, nil)
	c.JSON(http.StatusOK, successRes)
}
