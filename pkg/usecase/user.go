package usecase

import (
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/domain"
	helper_interfaces "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	usecase "CrocsClub/pkg/usecase/interfaces"
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
	wallet              interfaces.WalletRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, wallet interfaces.WalletRepository, otp interfaces.OtpRepository, inv interfaces.InventoryRepository, order interfaces.OrderRepository, h helper_interfaces.Helper) usecase.UserUseCase {
	return &userUseCase{
		userRepo:            repo,
		cfg:                 cfg,
		otpRepository:       otp,
		inventoryRepository: inv,
		orderRepository:     order,
		helper:              h,
		wallet:              wallet,
	}
}

var InternalError = "Internal Server Error"

func (u *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	if user.Name == "" {
		return models.TokenUsers{}, errors.New("username cannot be empty")
	}
	namevalidate, err := u.helper.ValidateAlphabets(user.Name)
	if err != nil {
		return models.TokenUsers{}, errors.New("invalid format for name")
	}
	if !namevalidate {
		return models.TokenUsers{}, errors.New("not a string")
	}

	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}

	phonenumber := u.helper.ValidatePhoneNumber(user.Phone)
	if !phonenumber {
		return models.TokenUsers{}, errors.New("invalid phone")
	}

	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}
	if user.Password == "" {
		return models.TokenUsers{}, errors.New("password cannot be empty")
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
	fmt.Println("user......", userDetails.Id)
	if err = u.wallet.CreateWallet(userDetails.Id); err != nil {
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

	products, err := u.userRepo.GetProductsInCart(cart_id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}

	var product_names []string
	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		product_names = append(product_names, product_name)
	}

	var quantity []int
	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		quantity = append(quantity, q)
	}

	var price []float64
	for i := range products {
		q, err := u.userRepo.FindPrice(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		price = append(price, q)
	}

	var categories []int
	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		categories = append(categories, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ID = products[i]
		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Price = int(price[i])
		get.Total = (price[i]) * float64(quantity[i])

		getcart = append(getcart, get)
	}

	var response models.GetCartResponse
	response.ID = cart_id
	response.Data = getcart

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
	if address.Name == "" || address.HouseName == "" || address.Street == "" || address.City == "" || address.State == "" || address.Phone == "" || address.Pin == "" {
		return errors.New("field cannot be empty")
	}
	ok, err := u.helper.ValidateAlphabets(address.Name)
	if err != nil {
		return errors.New("invalid format for name")
	}
	if !ok {
		return errors.New("invalid format for name")
	}
	phonenumber := u.helper.ValidatePhoneNumber(address.Phone)
	if !phonenumber {
		return errors.New("invalid phone")
	}
	pinnumber := u.helper.ValidatePin(address.Pin)
	if !pinnumber {
		return errors.New("invalid pin number")
	}

	rslt := u.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err = u.userRepo.AddAddress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}
	return nil
}

func (u *userUseCase) GetAddress(id int) ([]domain.Address, error) {

	address, err := u.userRepo.GetAddress(id)
	if err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return address, nil

}

func (u *userUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	details, err := u.userRepo.GetUserDetails(id)

	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}
	return details, nil
}

func (u *userUseCase) Edit(id int, user models.Edit) (models.Edit, error) {

	if user.Name == "" {
		return models.Edit{}, errors.New("name cannot be empty")
	}
	ok, err := u.helper.ValidateAlphabets(user.Name)
	if err != nil {
		return models.Edit{}, errors.New("invalid format for name")
	}
	if !ok {
		return models.Edit{}, errors.New("invalid format for name")
	}
	// namevalidate, err := u.helper.ValidateDatatype(user.Name, "string")
	// if err != nil {
	// 	return models.Edit{}, errors.New("invalid format for name")
	// }
	// if !namevalidate {
	// 	return models.Edit{}, errors.New("not a string")
	// }

	phonenumber := u.helper.ValidatePhoneNumber(user.Phone)
	if !phonenumber {
		return models.Edit{}, errors.New("invalid phone")
	}

	result, err := u.userRepo.Edit(id, user)

	if err != nil {
		return models.Edit{}, err
	}

	return result, nil
}

func (i *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return errors.New(InternalError)
	}

	err = i.helper.CompareHashAndPassword(userPassword, old)
	if err != nil {
		return errors.New("password incorrect")
	}
	fmt.Println("password", password)
	fmt.Println("repassword", repassword)
	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := i.helper.PasswordHashing(password)
	if err != nil {
		return errors.New("error in hashing password")
	}

	return i.userRepo.ChangePassword(id, string(newpassword))

}
