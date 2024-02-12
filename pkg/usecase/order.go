package usecase

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"strconv"
	"time"

	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
	cartRepo        interfaces.CartRepository
	couponRepo      interfaces.CouponRepository
	wallet          interfaces.WalletRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, wallet interfaces.WalletRepository, userUseCase services.UserUseCase, cartRepo interfaces.CartRepository, couponRepo interfaces.CouponRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
		cartRepo:        cartRepo,
		couponRepo:      couponRepo,
		wallet:          wallet,
	}
}

func (i *orderUseCase) OrderItemsFromCart(userID, addressID, paymentID, couponId int, useWallet bool) (models.OrderDetailsRep, error) {
	cart, err := i.userUseCase.GetCart(userID)
	var FinalPrice float64
	if err != nil {
		return models.OrderDetailsRep{}, err
	}
	exist, err := i.cartRepo.CheckCart(userID)

	if err != nil {
		return models.OrderDetailsRep{}, err
	}
	if !exist {
		return models.OrderDetailsRep{}, errors.New("cart is empty")
	}
	var total float64
	for _, item := range cart.Data {
		if item.Quantity > 0 && item.Price > 0 {
			total += float64(item.Quantity) * float64(item.Price)
		}
	}

	walletAmount, err := i.wallet.GetWallet(userID)
	if err != nil {
		return models.OrderDetailsRep{}, err
	}
	walletID, err := i.wallet.GetWalletId(userID)
	if err != nil {
		return models.OrderDetailsRep{}, err
	}

	// if coupon applied
	// var orderID int
	if couponId == 0 {
		var Totalamt float64

		if total < walletAmount.Amount {
			Totalamt = 0.0
			err = i.orderRepository.DebitWallet(userID, total)
			if err != nil {
				return models.OrderDetailsRep{}, err
			}

		} else {
			Totalamt = total - walletAmount.Amount
			err = i.orderRepository.DebitWallet(userID, walletAmount.Amount)
			if err != nil {
				return models.OrderDetailsRep{}, err
			}
		}

		orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, Totalamt)
		if err != nil {
			return models.OrderDetailsRep{}, err
		}
		err = i.wallet.WalletDebited(walletID, orderID, total)
		if err != nil {
			return models.OrderDetailsRep{}, err
		}

		if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
			return models.OrderDetailsRep{}, err
		}
	}
	if couponId != 0 {

		couponIdExist, err := i.couponRepo.CheckCouponById(couponId)
		if err != nil {
			return models.OrderDetailsRep{}, err
		}
		if !couponIdExist {
			return models.OrderDetailsRep{}, errors.New("coupon does not exist")
		}
		if couponId < 0 {
			return models.OrderDetailsRep{}, errors.New("negative values are not accepted")
		}
		coupondetails, err := i.couponRepo.GetCouponById(couponId)
		if err != nil {
			return models.OrderDetailsRep{}, errors.New("error in getting coupon")
		}

		finalprice := total - ((total * float64(coupondetails.DiscountPercentage)) / 100)
		var Totalamt float64
		if finalprice < walletAmount.Amount {
			Totalamt = 0.0
			err = i.orderRepository.DebitWallet(userID, finalprice)
			if err != nil {
				return models.OrderDetailsRep{}, err
			}

		} else {
			Totalamt = finalprice - walletAmount.Amount
			err = i.orderRepository.DebitWallet(userID, walletAmount.Amount)
			if err != nil {
				return models.OrderDetailsRep{}, err
			}
		}
		FinalPrice = finalprice
		orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, Totalamt)
		if err != nil {
			return models.OrderDetailsRep{}, err
		}
		err = i.wallet.WalletDebited(walletID, orderID, finalprice)
		if err != nil {
			return models.OrderDetailsRep{}, err
		}
		if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
			return models.OrderDetailsRep{}, err
		}

	}

	// last step
	for _, v := range cart.Data {
		if err := i.orderRepository.ReduceInventoryQuantity(v.ProductName, v.Quantity); err != nil {
			return models.OrderDetailsRep{}, err
		}
	}

	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return models.OrderDetailsRep{}, err
		}
	}
	var Order models.OrderDetailsRep

	Order.Total = int(total)
	Order.ID = cart.ID
	Order.PaymentMethod = paymentID
	Order.FinalPrice = FinalPrice

	return Order, nil
}

func (i *orderUseCase) GetOrders(orderId int) (domain.OrderResponse, error) {
	if orderId <= 0 {
		return domain.OrderResponse{}, errors.New("order ID should be a positive number")
	}

	orders, err := i.orderRepository.GetOrders(orderId)
	if err != nil {
		return domain.OrderResponse{}, err
	}
	return orders, nil
}

func (i *orderUseCase) CancelOrder(orderID int) error {
	orderStatus, err := i.orderRepository.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}

	if orderStatus != "PENDING" {
		return errors.New("order cannot be canceled, kindly return the product if accidentally booked")
	}

	if orderStatus == "CANCELED" {
		return errors.New("order cannot be canceled")
	}
	if orderStatus == "DELIVERED" {
		return errors.New("order cannot be canceled, kindly return the product if accidentally booked")
	}

	cart, err := i.orderRepository.GetOrders(orderID)
	if err != nil {
		return err
	}
	if cart.PaymentMethodID == 3 {
		ok, err := i.wallet.IsWalletExist(int(cart.UserID))
		if err != nil {
			return err
		}
		if !ok {
			err := i.wallet.CreateWallet(int(cart.UserID))
			if err != nil {
				return err
			}
		}
	}

	err = i.orderRepository.CancelOrder(orderID, int(cart.UserID), int(cart.FinalPrice), cart.PaymentStatus)
	if err != nil {
		return err
	}

	walletID, err := i.wallet.GetWalletId(int(cart.UserID))
	if err != nil {
		return err
	}

	err = i.wallet.WalletCredited(walletID, orderID, cart.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (i *orderUseCase) GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {
	allorder, err := i.orderRepository.GetAllOrders(userId, page, pageSize)
	if err != nil {
		return []models.OrderDetails{}, err
	}
	return allorder, nil
}

func (i *orderUseCase) GetAdminOrders(page int) ([]models.CombinedOrderDetails, error) {

	orderDetails, err := i.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil
}

func (i *orderUseCase) OrdersStatus(orderID string) error {
	status, err := i.orderRepository.CheckOrdersStatusByID(orderID)
	if err != nil {
		return err
	}

	switch status {
	case "CANCELED", "RETURNED", "DELIVERED":
		return errors.New("cannot approve this order because it's in a processed or canceled state")
	case "PENDING":
		err := i.orderRepository.ChangeOrderStatus(orderID, "SHIPPED")
		if err != nil {
			return err
		}
	case "SHIPPED":
		shipmentStatus, err := i.orderRepository.GetShipmentStatus(orderID)
		if err != nil {
			return err
		}

		if shipmentStatus == "CANCELLED" {
			return errors.New("cannot approve this order because it's cancelled")
		}

		err = i.orderRepository.ChangeOrderStatus(orderID, "DELIVERED")
		if err != nil {
			return err
		}
	}

	return nil
}

func (or *orderUseCase) PaymentMethodID(order_id int) (int, error) {
	id, err := or.orderRepository.PaymentMethodID(order_id)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	fmt.Println("order inside order usecase", id)
	return id, nil
}

func (o *orderUseCase) ReturnOrder(orderID string) error {
	var ReturnOrderResponse models.ReturnOrderResponse

	orderIDint, err := strconv.Atoi(orderID)
	if err != nil {
		return errors.New("string convertion failed")
	}
	cart, err := o.orderRepository.GetOrders(orderIDint)
	if err != nil {
		return err
	}
	ReturnOrderResponse.CartAmount = int(cart.FinalPrice)
	ReturnOrderResponse.UserId = int(cart.UserID)
	ReturnOrderResponse.OrderStatus = "Return To Wallet"
	ReturnOrderResponse.OrderID = orderIDint
	if cart.PaymentMethodID == 3 || cart.PaymentMethodID == 1 {
		ok, err := o.wallet.IsWalletExist(int(cart.UserID))
		if err != nil {
			return err
		}
		if !ok {

			err := o.wallet.CreateWallet(int(cart.UserID))
			if err != nil {
				return err
			}
		}
	}

	shipmentStatus, err := o.orderRepository.GetShipmentsStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "DELIVERED" {

		b := o.orderRepository.ReturnOrder(ReturnOrderResponse)
		walletID, err := o.wallet.GetWalletId(int(cart.UserID))

		if err != nil {
			return err
		}
		err = o.wallet.WalletCredited(walletID, orderIDint, cart.FinalPrice)
		if err != nil {
			return err
		}
		return b
	}

	return errors.New("can't return order")

}

func (or *orderUseCase) PrintInvoice(orderId int) (*gofpdf.Fpdf, error) {

	if orderId < 1 {
		return nil, errors.New("enter a valid order id")
	}

	order, err := or.orderRepository.GetDetailedOrderThroughId(orderId)
	if err != nil {
		return nil, err
	}

	fmt.Println("order usecase ", order)

	items, err := or.orderRepository.GetItemsByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	fmt.Println("items usecase", items)

	if order.OrderStatus != "DELIVERED" {
		return nil, errors.New("wait for the invoice until the product is received")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 30)
	pdf.SetTextColor(31, 73, 125)
	pdf.Cell(0, 20, "Invoice")
	pdf.Ln(20)

	pdf.SetFont("Arial", "I", 14)
	pdf.SetTextColor(51, 51, 51)
	pdf.Cell(0, 10, "Customer Details")
	pdf.Ln(10)
	customerDetails := []string{
		"Name: " + order.Name,
		"House Name: " + order.HouseName,
		"Street: " + order.Street,
		"State: " + order.State,
		"City: " + order.City,
	}
	for _, detail := range customerDetails {
		pdf.Cell(0, 10, detail)
		pdf.Ln(10)
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 10, "Item", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Price", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Final Price", "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255)
	for _, item := range items {
		pdf.CellFormat(40, 10, item.ProductName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price, 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, strconv.Itoa(item.Quantity), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price*float64(item.Quantity), 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(10)

	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Total Price:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(totalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	offerApplied := totalPrice - order.FinalPrice

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Offer Applied:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(offerApplied, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Final Amount:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(order.FinalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, "Generated by Crocsclub India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))
	pdf.Ln(10)

	return pdf, nil
}
