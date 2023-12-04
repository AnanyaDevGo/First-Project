package usecase

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"

	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"
	"log"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
	cartRepo        interfaces.CartRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase, cartRepo interfaces.CartRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
		cartRepo:        cartRepo,
	}
}

func (i *orderUseCase) OrderItemsFromCart(userID, addressID, paymentID int) error {
	cart, err := i.userUseCase.GetCart(userID)
	if err != nil {
		return err
	}
	exist, err := i.cartRepo.CheckCart(userID)

	if err != nil {
		return err
	}

	if !exist {
		return errors.New("cart is empty")
	}
	var total float64
	for _, item := range cart.Data {
		if item.Quantity > 0 && item.Price > 0 {
			total += float64(item.Quantity) * float64(item.Price)
		}
	}

	fmt.Println("total at usecase", total)

	orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, total)
	if err != nil {
		return err
	}
	if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
		return err
	}
	for _, v := range cart.Data {
		if err := i.orderRepository.ReduceInventoryQuantity(v.ProductName, v.Quantity); err != nil {
			return err
		}
	}

	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}

	return nil
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
		return errors.New("order cannot be canceled")
	}

	cart, err := i.orderRepository.GetOrders(orderID)
	if err != nil {
		return err
	}

	err = i.orderRepository.CancelOrder(orderID, int(cart.UserID), int(cart.FinalPrice), cart.PaymentStatus)
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
	fmt.Println("here order")
	id, err := or.orderRepository.PaymentMethodID(order_id)
	// fmt.Println("id", id, err)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	fmt.Println("order inside order usecase", id)
	return id, nil
}

func (o *orderUseCase) ReturnOrder(orderID string) error {

	shipmentStatus, err := o.orderRepository.GetShipmentsStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "DELIVERED" {
		shipmentStatus = "RETURNED"
		return o.orderRepository.ReturnOrder(shipmentStatus, orderID)
	}

	return errors.New("can't return order")

}
