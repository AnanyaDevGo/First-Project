package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
	GetCart(id int) (models.GetCartResponse, error)
	RemoveFromCart(cart, inventory int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	GetAddress(id int) ([]domain.Address, error)
	AddAddress(id int, address models.AddAddress) error
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
	ChangePassword(id int, old string, password string, repassword string) error
}
