package usecase

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"errors"
	"fmt"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
	}
}

func (o *orderUseCase) GetOrders(id int) ([]domain.Order, error) {
	order, err := o.orderRepository.GetOrders(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("order", order)
	return order, nil
}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	returned, err := i.orderRepository.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Returned = returned
	response.Delivered = delivered

	return response, nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	status, err := i.orderRepository.CheckOrderStatusByID(id)
	if err != nil {
		return err
	}

	if status != "PENDING" {
		return errors.New("order cannot be canceled if you accidently booked kindly return the product")
	}

	err = i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}
