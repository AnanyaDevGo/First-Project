package interfaces

import "CrocsClub/pkg/utils/models"



type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
}

