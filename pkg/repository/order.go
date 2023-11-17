package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
	"fmt"

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
	fmt.Println("order repo", order)
	return order, nil
}

func (i *orderRepository) OrderItems(userid, addressid, paymentid int, total float64, coupon string) (int, error) {

	var id int
	query := `
    INSERT INTO orders (user_id,address_id, payment_method_id, final_price,coupon_used)
    VALUES (?, ?, ?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, userid, addressid, paymentid, total, coupon).Scan(&id)

	return id, nil

}

func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `
    INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `

	for _, v := range cart {
		var inv int
		if err := i.DB.Raw("select id from inventories where product_name=$1", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}

		if err := i.DB.Exec(query, order_id, inv, v.Quantity).Error; err != nil {
			return err
		}
	}

	return nil

}

func (i *orderRepository) CancelOrder(id int) error {

	if err := i.DB.Exec("update orders set order_status='CANCELED' where id=$1", id).Error; err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) EditOrderStatus(status string, id int) error {

	if err := o.DB.Exec("update orders set order_status=$1 where id=$2", status, id).Error; err != nil {
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
func (or *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {

	var orders []domain.OrderDetails
	if err := or.DB.Raw("SELECT orders.id AS id, users.name AS username, CONCAT('House Name:',addresses.house_name, ',', 'Street:', addresses.street, ',', 'City:', addresses.city, ',', 'State', addresses.state, ',', 'Phone:', addresses.phone) AS address, payment_methods.payment_name AS payment_method, orders.final_price As total FROM orders JOIN users ON users.id = orders.user_id JOIN payment_methods ON payment_methods.id = orders.payment_method_id JOIN addresses ON orders.address_id = addresses.id WHERE order_status = $1", status).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}

	return orders, nil

}

func (o *orderRepository) GetOrderDetail(orderID string) (domain.Order, error) {

	var orderDetails domain.Order
	err := o.DB.Raw("select * from orders where order_id = ?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return domain.Order{}, err
	}

	return orderDetails, nil

}

func (o *orderRepository) MakePaymentStatusAsPaid(id int) error {

	err := o.DB.Exec("UPDATE orders SET payment_status = 'PAID' WHERE id = $1", id).Error
	if err != nil {
		return err
	}

	return nil
}
