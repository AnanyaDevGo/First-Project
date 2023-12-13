package usecase

import (
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUseCase         services.UserUseCase
	adrepo              interfaces.AdminRepository
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase, adrepo interfaces.AdminRepository) services.CartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUseCase:         userUseCase,
		adrepo:              adrepo,
	}
}

func (i *cartUseCase) AddToCart(userID, inventoryID, qty int) error {

	fmt.Println("uuuuuuuuuu", inventoryID)

	stock, err := i.inventoryRepository.CheckStock(inventoryID)
	if err != nil {
		return err
	}

	fmt.Println("stock", stock)

	if stock <= 0 || qty > stock {
		return errors.New("out of stock")
	}

	cart_id, err := i.repo.GetCartId(userID)
	if err != nil {
		return errors.New("some error in geting user cart")
	}
	fmt.Println("cart idjjjjjg", cart_id)
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(userID)
		if err != nil {
			return errors.New("cannot create cart fro user")
		}
	}

	exists, err := i.repo.CheckIfItemIsAlreadyAdded(cart_id, inventoryID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("item already exists in cart")
	}

	if err := i.repo.AddLineItems(cart_id, inventoryID, qty); err != nil {
		return errors.New("error in adding products")
	}

	

	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	if id <= 0 {
		return models.CheckOut{}, errors.New("invalid id")
	}
	fmt.Println("cart id .........", id)
	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymethods, err := i.adrepo.GetPaymentMethod()
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.userUseCase.GetCart(id)
	if err != nil {
		return models.CheckOut{}, err
	}
	var checkout models.CheckOut

	checkout.CartID = id
	checkout.Addresses = address
	checkout.Products = products.Data
	checkout.PaymentMethod = paymethods

	// fmt.Println("cartid", checkout.CartID)
	// fmt.Println("cartaddress", checkout.Addresses)
	// fmt.Println("cartproducts", checkout.Products)
	fmt.Println("cartpay", checkout.PaymentMethod)

	return checkout, err
}
