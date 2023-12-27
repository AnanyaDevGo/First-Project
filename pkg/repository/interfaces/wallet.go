package interfaces

import "CrocsClub/pkg/utils/models"

type WalletRepository interface {
	CreateWallet(userID int) error
	GetWallet(userID int) (models.WalletAmount, error)
	WalletHistory(walletId int) ([]models.WalletHistory, error)
	IsWalletExist(userID int) (bool, error)
	WalletCredited(walletID, OrderID int, Amount float64) error
	WalletDebited(walletID, OrderID int, Amount float64) error
	GetWalletId(userID int) (int, error)
}
