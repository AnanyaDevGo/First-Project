package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
	"errors"
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
func (i *orderRepository) GetOrders(orderID int) (domain.OrderResponse, error) {
	if orderID <= 0 {
		return domain.OrderResponse{}, errors.New("order ID should'n be a negative number")
	}

	var order domain.OrderResponse

	query := `SELECT * FROM orders WHERE id = $1`

	if err := i.DB.Raw(query, orderID).First(&order).Error; err != nil {
		return domain.OrderResponse{}, err
	}

	return order, nil
}

func (ad *orderRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=$1", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}
	return cart, nil

}
func (i *orderRepository) OrderItems(userid, addressid, paymentid int, total float64) (int, error) {

	var id int
	query := `
    INSERT INTO orders (created_at,user_id,address_id, payment_method_id, final_price)
    VALUES (Now(),?, ?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, userid, addressid, paymentid, total).Scan(&id)
	fmt.Println("id...........", id)
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

		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
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

func (i *orderRepository) GetAllOrders(userID, page, pageSize int) ([]models.OrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var order []models.OrderDetails

	err := i.DB.Raw("SELECT id as order_id, address_id, payment_method_id, final_price as price, order_status, payment_status FROM orders WHERE user_id = ? OFFSET ? LIMIT ?", userID, offset, pageSize).Scan(&order).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Retrieved orders:", order)
	return order, nil
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

func (o *orderRepository) FindPaymentMethodOfOrder(id int) (string, error) {

	var payment string

	if err := o.DB.Raw(`select payment_methods.payment_name
	 from payment_methods
	  join orders on orders.payment_method_id = payment_methods.id
	   where orders.id = $1`, id).Scan(&payment).Error; err != nil {
		return "", err
	}
	return payment, nil
}

func (o *orderRepository) GetProductImagesInAOrder(id int) ([]string, error) {

	var images []string
	err := o.DB.Raw(`SELECT inventories.image
	FROM order_items 
	JOIN inventories ON inventories.id = order_items.inventory_id
	JOIN orders ON orders.id = order_items.order_id 
	WHERE orders.id = $1`, id).Scan(&images).Error
	if err != nil {
		return []string{}, err
	}

	return images, nil
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
