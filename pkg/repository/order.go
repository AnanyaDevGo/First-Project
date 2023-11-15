package repository

import (
	"CrocsClub/pkg/domain"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository) GetOrders(id int) ([]domain.Order, error) {
	var order []domain.Order
	if err := o.DB.Raw("select * from orders where user_id=?", id).Scan(&order).Error; err != nil {
		return []domain.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) EditOrderStatus(status string, id int) error {

	if err := o.DB.Exec("update orders set order_status=$1 where id=$2", status, id).Error; err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) CancelOrder(id int) error {

	if err := o.DB.Exec("update orders set order_status='CANCELED' where id=$1", id).Error; err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) CheckOrderStatusByID(id int) (string, error) {

	var status string
	err := o.DB.Raw("select order_status from orders where id = ?", id).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}
