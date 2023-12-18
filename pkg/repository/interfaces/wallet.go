package interfaces

import "CrocsClub/pkg/utils/models"

type WalletRepository interface {
	CreateWallet(userID int) error
	GetWallet(userID int) (models.WalletAmount, error)
	WalletHistory(userID int) ([]models.WalletHistory, error)
	IsWalletExist(userID int) (bool, error)
}
