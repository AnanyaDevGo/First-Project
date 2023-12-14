package usecase

import (
	"CrocsClub/pkg/config"
	helper_interfaces "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
	helper        helper_interfaces.Helper
}

func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository, h helper_interfaces.Helper) services.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
		helper:        h,
	}
}

func (ot *otpUseCase) SendOTP(phone string) error {

	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	phonenumber := ot.helper.ValidatePhoneNumber(phone)
	if !phonenumber {
		return errors.New("invalid phone number")
	}

	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	fmt.Println("accsid:", ot.cfg.SERVICESID)

	_, err := ot.helper.TwilioSendOTP(phone, ot.cfg.SERVICESID)
	if err != nil {
		return errors.New("error ocurred while generating OTP" + err.Error())
	}

	return nil

}

func (ot *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {

	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := ot.helper.TwilioVerifyOTP(ot.cfg.SERVICESID, code.Code, code.PhoneNumber)
	if err != nil {

		return models.TokenUsers{}, errors.New("error while verifying")
	}

	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := ot.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	var user models.UserDetailsResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil

}
