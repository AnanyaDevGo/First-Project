package usecase

import (
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
)

type walletUseCase struct {
	walletRepository interfaces.WalletRepository
}

func NewWalletUseCase(repo interfaces.WalletRepository) services.WalletUseCase {
	return &walletUseCase{
		walletRepository: repo,
	}
}
func (w *walletUseCase) GetWallet(userID int) (models.WalletAmount, error) {
	return w.walletRepository.GetWallet(userID)
}
func (w *walletUseCase) WalletHistory(userID int) ([]models.WalletHistory, error) {
	walletId, err := w.walletRepository.GetWalletId(userID)
	if err != nil {
		return []models.WalletHistory{}, err
	}

	return w.walletRepository.WalletHistory(walletId)
}
