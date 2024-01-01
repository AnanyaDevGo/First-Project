package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userDataBase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDataBase{DB}
}

func (c *userDataBase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (c *userDataBase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse

	err := c.DB.Raw("INSERT INTO users (name, email, password, phone) VALUES (?, ?, ?, ?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, fmt.Errorf("failed to sign up user: %w", err)
	}

	userId := userDetails.Id

	err = c.DB.Raw("INSERT INTO wallets (user_id, amount) VALUES (?, '0')", userId).Error
	if err != nil {

		return models.UserDetailsResponse{}, fmt.Errorf("failed to create wallet: %w", err)
	}

	return userDetails, nil
}

func (cr *userDataBase) UserBlockStatus(email string) (bool, error) {
	fmt.Println(email)
	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}
	fmt.Println(isBlocked)
	return isBlocked, nil
}
func (c *userDataBase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var user_details models.UserSignInResponse

	err := c.DB.Raw(`
		SELECT * FROM users where email = ? and blocked = false
		`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return user_details, nil

}

func (ad *userDataBase) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("select inventories.product_name,cart_products.quantity,cart_products.total_price from cart_products inner join inventories on cart_products.inventory_id=inventories.id where user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}

func (ad *userDataBase) GetCartID(id int) (int, error) {

	var cart_id int

	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id = ?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}

	return cart_id, nil

}

func (ad *userDataBase) FindCartQuantity(cart_id, inventory_id int) (int, error) {

	var quantity int

	if err := ad.DB.Raw("select quantity from line_items where cart_id=$1 and inventory_id=$2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}

	return quantity, nil

}

func (ad *userDataBase) FindPrice(inventory_id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from inventories where id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil

}

func (ad *userDataBase) FindCategory(inventory_id int) (int, error) {

	var category int

	if err := ad.DB.Raw("select category_id from inventories where id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}

	return category, nil

}

func (i *userDataBase) FindStock(id int) (int, error) {
	var stock int
	err := i.DB.Raw("SELECT stock FROM inventories WHERE id = ?", id).Scan(&stock).Error
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (ad *userDataBase) FindProductNames(inventory_id int) (string, error) {

	var product_name string

	if err := ad.DB.Raw("select product_name from inventories where id=?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}

	return product_name, nil

}

func (ad *userDataBase) GetProductsInCart(cart_id int) ([]int, error) {

	var cart_products []int

	if err := ad.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}

	return cart_products, nil

}

func (ad *userDataBase) RemoveFromCart(cart, inventory int) error {

	if err := ad.DB.Exec(`DELETE FROM line_items WHERE cart_id = $1 AND inventory_id = $2`, cart, inventory).Error; err != nil {
		return err
	}
	return nil
}
func (ad *userDataBase) UpdateQuantity(id, invID, qty int) error {

	if id <= 0 || invID <= 0 || qty <= 0 {
		return errors.New("negative or zero values are not allowed")
	}

	if qty >= 20 {
		return errors.New("choose number of items below 20")
	}

	if qty >= 0 {
		query := `
			UPDATE line_items
			SET quantity = $1
			WHERE cart_id = $2 AND inventory_id = $3
		`

		result := ad.DB.Exec(query, qty, id, invID)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (ad *userDataBase) AddAddress(id int, address models.AddAddress, result bool) error {
	err := ad.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, phone, pin,"default")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 )`,
		id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Phone, address.Pin, result).Error
	if err != nil {
		return errors.New("could not add address")
	}

	return nil
}

func (c *userDataBase) CheckIfFirstAddress(id int) bool {

	var count int
	if err := c.DB.Raw("select count(*) from addresses where user_id=$1", id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0

}

func (ad *userDataBase) GetAddress(id int) ([]domain.Address, error) {
	var address []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&address).Error; err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return address, nil

}

func (ad *userDataBase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	var details models.UserDetailsResponse
	if err := ad.DB.Raw("select id,name,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}
	return details, nil
}

func (ad *userDataBase) Edit(id int, user models.Edit) (models.Edit, error) {
	var result models.Edit

	args := []interface{}{}
	query := "update users set"

	if user.Email != "" {
		query += " email = $1,"

		args = append(args, user.Email)
	}

	if user.Name != "" {
		query += " name = $2,"
		args = append(args, user.Name)
	}

	if user.Phone != "" {
		query += " phone = $3,"

		args = append(args, user.Phone)
	}

	query = query[:len(query)-1] + " where id = $4"

	args = append(args, id)

	err := ad.DB.Exec(query, args...).Error
	if err != nil {
		return models.Edit{}, err
	}
	query2 := "select * from users where id = ?"
	if err := ad.DB.Raw(query2, id).Scan(&result).Error; err != nil {
		return models.Edit{}, err
	}

	return result, nil

}

func (ad *userDataBase) ChangePassword(id int, password string) error {

	if id <= 0 {
		return errors.New("negative or zero values are not allowed")
	}

	err := ad.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
		return err
	}

	return nil

}

func (u *userDataBase) GetPassword(id int) (string, error) {

	if id <= 0 {
		return "", errors.New("negative or zero values are not allowed")
	}

	var userPassword string
	err := u.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}
