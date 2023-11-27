package interfaces

import "CrocsClub/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, inventory_id, qty int) error
	CheckOut(id int) (models.CheckOut, error)
}
