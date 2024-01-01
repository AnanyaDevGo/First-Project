package handler

import (
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// @Summary Send OTP
// @Description Send OTP to the specified phone number for verification.
// @Accept json
// @Produce json
// @Tags User
// @Param body body models.OTPData true "Phone number details in JSON format"
// @Success 200 {object} response.Response "OTP sent successfully"
// @Failure 400 {object} response.Response "Fields provided in wrong format or Could not send OTP"
// @Router /user/otplogin [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Verify OTP
// @Description Verify the provided OTP code for user authentication.
// @Accept json
// @Produce json
// @Tags User
// @Param body body models.VerifyData true "Verification code details in JSON format"
// @Success 200 {object} response.Response "Successfully verified OTP"
// @Failure 400 {object} response.Response "Fields provided in wrong format or Could not verify OTP"
// @Router /user/verifyotp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	users, err := ot.otpUseCase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusAccepted, successRes)

}
