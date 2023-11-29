package interfaces

import "CrocsClub/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error)
	VerifyPayment(paymentID string, razorID string, orderID string) error
}
