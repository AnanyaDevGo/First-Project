package repository

import (
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (ad *cartRepository) GetAddresses(id int) ([]models.Address, error) {

	var addresses []models.Address

	if err := ad.DB.Raw("SELECT * FROM addresses WHERE user_id=$1", id).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}

	return addresses, nil

}

func (ad *cartRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=$1", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}

func (ad *cartRepository) GetCartId(user_id int) (int, error) {

	var id int

	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (i *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	err := i.DB.Exec(`
		INSERT INTO carts (user_id)
		VALUES ($1)`, user_id).Error
	if err != nil {
		return 0, err
	}

	if err := i.DB.Raw("select id from carts where user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (i *cartRepository) AddLineItems(cart_id, inventory_id, qty int) error {

	err := i.DB.Exec(`
		INSERT INTO line_items (cart_id,inventory_id,quantity)
		VALUES ($1,$2,$3)`, cart_id, inventory_id, qty).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *cartRepository) CheckIfItemIsAlreadyAdded(cart_id, inventory_id int) (bool, error) {

	var count int

	if err := ad.DB.Raw("SELECT COUNT(*) FROM line_items WHERE cart_id = $1 AND inventory_id = $2", cart_id, inventory_id).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil

}

func (cr *cartRepository) CheckCart(userID int) (bool, error) {
	var count int
	querry := `	SELECT COUNT(*) 
	FROM line_items 
	WHERE cart_id IN (SELECT id FROM carts WHERE user_id = ?)
	`
	err := cr.DB.Raw(querry, userID).Scan(&count).Error
	if err != nil {
		return false, errors.New("no cart found")
	}
	if count <= 0 {
		return false, errors.New("no cart found")
	}
	return true, nil
}

// func (cr *cartRepository) GetTotalPriceFromCart(cartId int)(int, error){
// 	err := cr.DB.Raw("select ", userID).Scan(&count).Error
// 	if err != nil {
// 		return false, errors.New("no cart found")
// 	}
// }
