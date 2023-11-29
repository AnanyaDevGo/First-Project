package usecase

import (
	interfaces "CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"fmt"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUsecase struct {
	repository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUsecase{
		repository: repo,
	}
}

func (p *paymentUsecase) MakePaymentRazorPay(orderID string, newuserid int) (models.OrderPaymentDetails, error) {
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	//get userid
	// newuserid, err := strconv.Atoi(userID)
	// if err != nil {
	// 	return models.OrderPaymentDetails{}, err
	// }

	orderDetails.UserID = newuserid

	//get username
	username, err := p.repository.FindUsername(newuserid)
	fmt.Println("username", username)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	//get total
	newfinal, err := p.repository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.FinalPrice = newfinal

	client := razorpay.NewClient("rzp_test_dOkReCdGayq2PC", "8cob7I0S3BEywzFKlkTnMQiy")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	fmt.Println("body", body)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}

	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}

func (p *paymentUsecase) VerifyPayment(paymentID string, razorID string, orderID string) error {
	fmt.Println("paymetn Id", paymentID, " razorId ", razorID, " orderId ", orderID)

	err := p.repository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	return nil

}
