package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (i *orderRepository) GetOrders(orderID int) (domain.OrderResponse, error) {

	var order domain.OrderResponse

	query := `SELECT * FROM orders WHERE id = $1`

	if err := i.DB.Raw(query, orderID).First(&order).Error; err != nil {
		return domain.OrderResponse{}, err
	}
	fmt.Println("rrrrrr.....", order.FinalPrice)

	return order, nil
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

func (o *orderRepository) CheckOrderStatusByID(id int) (string, error) {

	var status string
	err := o.DB.Raw("select order_status from orders where id = ?", id).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}
func (i *orderRepository) ReduceInventoryQuantity(productName string, quantity int) error {
	query := `
        UPDATE inventories
        SET stock = stock - ?
        WHERE product_name = ?
    `
	if err := i.DB.Exec(query, quantity, productName).Error; err != nil {
		return err
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

func (o *orderRepository) GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2

	var orderDetails []models.CombinedOrderDetails

	err := o.DB.Raw(`
	SELECT orders.id AS order_id, orders.final_price, orders.order_status, orders.payment_status, 
	users.name, users.email, users.phone, addresses.house_name, addresses.state, 
	addresses.pin, addresses.street, addresses.city 
	FROM orders 
	INNER JOIN users ON orders.user_id = users.id 
	INNER JOIN addresses ON users.id = addresses.user_id 
	LIMIT ? OFFSET ?
`, 2, offset).Scan(&orderDetails).Error

	if err != nil {
		return []models.CombinedOrderDetails{}, nil
	}

	return orderDetails, nil
}

func (o *orderRepository) CheckOrdersStatusByID(id string) (string, error) {
	var status string
	err := o.DB.Raw("SELECT order_status FROM orders WHERE id = ?", id).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func (i *orderRepository) GetShipmentStatus(orderID string) (string, error) {
	var shipmentStatus string
	err := i.DB.Exec("UPDATE orders SET order_status = 'DELIVERED', payment_status = 'PAID' WHERE id = ?", orderID).Error
	if err != nil {
		return "", err
	}
	return shipmentStatus, nil
}

func (or *orderRepository) GetOrderDetailsByOrderId(orderID int) (models.CombinedOrderDetails, error) {

	var orderDetails models.CombinedOrderDetails
	err := or.DB.Raw(`SELECT
    orders.id as order_id,
    orders.final_price,
    orders.shipment_status,
    orders.payment_status,
    users.name,
    users.email,
    users.phone,
    addresses.house_name,
    addresses.state,
    addresses.street,
    addresses.city,
    addresses.pin
FROM
    orders
INNER JOIN
    users ON orders.user_id = users.id
INNER JOIN
    addresses ON users.id = addresses.user_id
WHERE
    orders.id = ?`, orderID).Scan(&orderDetails).Error
	if err != nil {
		return models.CombinedOrderDetails{}, nil
	}
	fmt.Println("print", orderDetails.FinalPrice)

	return orderDetails, nil
}

func (i *orderRepository) ApproveOrder(orderID string) error {
	err := i.DB.Exec("UPDATE orders SET order_status = 'order_placed' WHERE id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *orderRepository) ChangeOrderStatus(orderID, status string) error {
	err := i.DB.Exec("UPDATE orders SET order_status = ? WHERE id = ?", status, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) GetShipmentsStatus(orderID string) (string, error) {

	var orderStatus string
	err := o.DB.Raw("select order_status from orders where id = ?", orderID).Scan(&orderStatus).Error
	if err != nil {
		return "", err
	}

	return orderStatus, nil

}
func (or *orderRepository) PaymentMethodID(orderID int) (int, error) {
	var a int
	err := or.DB.Raw("SELECT payment_method_id FROM orders WHERE id = ?", orderID).Scan(&a).Error
	if err != nil {
		return 0, err
	}
	fmt.Println("order id 1", a)
	return a, nil
}
func (o *orderRepository) ReturnOrder(orderStatus string, orderID string) error {

	err := o.DB.Exec("update orders set order_status = ? where id = ?", orderStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

func (or *orderRepository) PaymentAlreadyPaid(orderID int) (bool, error) {
	var a bool
	err := or.DB.Raw("SELECT order_status = 'processing' AND payment_status = 'paid' FROM orders WHERE id = ?", orderID).Scan(&a).Error
	if err != nil {
		return false, err
	}
	return a, nil
}
