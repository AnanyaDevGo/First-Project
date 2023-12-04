package repository

import (
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"

	"gorm.io/gorm"
)

type walletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) interfaces.WalletRepository {
	return &walletRepository{
		DB: db,
	}
}

func (w *walletRepository) GetWallet(userID int) (models.WalletAmount, error) {
	var walletAmount models.WalletAmount
	if err := w.DB.Raw("select amount from wallets where user_id = ?", userID).Scan(&walletAmount).Error; err != nil {
		return models.WalletAmount{}, err
	}
	return walletAmount, nil
}
func (w *walletRepository) WalletHistory(userID int) ([]models.WalletHistory, error) {
	var history []models.WalletHistory
	err := w.DB.Raw("SELECT id,order_id,description,amount,is_credited FROM wallet_histories WHERE user_id = ?", userID).Scan(&history).Error
	if err != nil {
		return []models.WalletHistory{}, err
	}
	return history, nil
}
