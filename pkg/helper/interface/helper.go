package interfaces

import (
	"CrocsClub/pkg/utils/models"
	"time"
)

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
	PasswordHashing(string) (string, error)
	CompareHashAndPassword(a string, b string) error
	Copy(a *models.UserDetailsResponse, b *models.UserSignInResponse) (models.UserDetailsResponse, error)
	GetTimeFromPeriod(timePeriod string) (time.Time, time.Time)
}
