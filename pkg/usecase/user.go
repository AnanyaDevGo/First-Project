package usecase

import (
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/domain"
	helper_interfaces "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo            interfaces.UserRepository
	cfg                 config.Config
	otpRepository       interfaces.OtpRepository
	inventoryRepository interfaces.InventoryRepository
	orderRepository     interfaces.OrderRepository
	helper              helper_interfaces.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, otp interfaces.OtpRepository, inv interfaces.InventoryRepository, order interfaces.OrderRepository, h helper_interfaces.Helper) *userUseCase {
	return &userUseCase{
		userRepo:            repo,
		cfg:                 cfg,
		otpRepository:       otp,
		inventoryRepository: inv,
		orderRepository:     order,
		helper:              h,
	}
}

var InternalError = "Internal Server Error"

func (u *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	fmt.Println("add users")

	userExist := u.userRepo.CheckUserAvailability(user.Email)
	fmt.Println("user exists", userExist)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}
	fmt.Println(user)
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := u.helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}

func (u *userUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}

	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (u *userUseCase) GetCart(id int) (models.GetCartResponse, error) {
	cart_id, err := u.userRepo.GetCartID(id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	fmt.Println("cart id ", cart_id)
	products, err := u.userRepo.GetProductsInCart(cart_id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}

	fmt.Println("products ", products)
	var product_names []string
	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		product_names = append(product_names, product_name)
	}
	fmt.Println("product names", product_names)

	var quantity []int
	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		quantity = append(quantity, q)
	}
	fmt.Println("quantity", quantity)

	var categories []int
	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		categories = append(categories, c)
	}
	fmt.Println("categories", categories)

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ID = products[i]
		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		getcart = append(getcart, get)
	}

	fmt.Println("get carts", getcart)
	var response models.GetCartResponse
	response.ID = cart_id
	response.Data = getcart

	fmt.Println("response", response)
	return response, nil

}

func (u *userUseCase) RemoveFromCart(cart, inventory int) error {

	err := u.userRepo.RemoveFromCart(cart, inventory)
	if err != nil {

		return err

	}
	return nil

}

func (i *userUseCase) UpdateQuantity(id, inv_id, qty int) error {

	err := i.userRepo.UpdateQuantity(id, inv_id, qty)
	if err != nil {
		return err
	}

	return nil

}

func (u *userUseCase) AddAddress(id int, address models.AddAddress) error {
	rslt := u.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := u.userRepo.AddAddress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}
	return nil
}

func (u *userUseCase) GetAddress(id int) ([]domain.Address, error) {

	address, err := u.userRepo.GetAddress(id)

	if err != nil {
		return []domain.Address{}, errors.New("error in getting address")
	}
	return address, nil
}

func (u *userUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	details, err := u.userRepo.GetUserDetails(id)

	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}
	return details, err
}

func (u *userUseCase) Edit(id int, user models.Edit) (models.Edit, error) {
	result, err := u.userRepo.Edit(id, user)
	if err != nil {
		return models.Edit{}, err
	}

	return result, nil
}

func (u *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {
	userPassword, err := u.userRepo.GetPassword(id)
	if err != nil {
		return errors.New(InternalError)
	}

	err = u.helper.CompareHashAndPassword(userPassword, old)
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := u.helper.PasswordHashing(password)
	if err != nil {
		return errors.New("error in hashing password")
	}

	return u.userRepo.ChangePassword(id, string(newpassword))
}
