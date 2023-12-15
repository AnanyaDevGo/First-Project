package interfaces

import (
	"CrocsClub/pkg/utils/models"
)

type CartRepository interface {
	GetCart(id int) ([]models.GetCart, error)
	GetAddresses(id int) ([]models.Address, error)
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	AddLineItems(cart_id, inventory_id, qty int) error
	CheckIfItemIsAlreadyAdded(cart_id, inventory_id int) (bool, error)
	CheckCart(userID int) (bool, error)
	// GetTotalPriceFromCart(cartId int) (int, error)
}
