package interfaces

import "CrocsClub/pkg/utils/models"

type WalletUseCase interface {
	GetWallet(userID int) (models.WalletAmount, error)
	WalletHistory(userID int) ([]models.WalletHistory, error)
}
