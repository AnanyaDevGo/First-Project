package usecase

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"errors"
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
	return order, nil
}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

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
