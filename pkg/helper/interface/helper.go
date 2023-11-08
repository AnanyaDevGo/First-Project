package interfaces

import "CrocsClub/pkg/utils/models"

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
}
