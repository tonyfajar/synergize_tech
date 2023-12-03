package repositories

import (
	"github.com/synergize_tech/db"
	"github.com/synergize_tech/models"
)

func CreateWallet(wallet *models.Wallet) error {

	return db.DB.Create(wallet).Error

}

func GetDetailsWallet(UserId uint) (models.Wallet, error) {
	var wallet models.Wallet
	db.DB.Where("user_id = ? ", UserId).First(&wallet)
	return wallet, nil
}

func TopUpWallet(UserId uint, amount float64) error {
	var wallet models.Wallet

	// Find the user by ID
	db.DB.Where("user_id = ? ", UserId).First(&wallet)
	// Update the age by incrementing it
	wallet.Balance += amount

	// Save the changes
	if err := db.DB.Save(&wallet).Error; err != nil {
		return err
	}

	return nil
}
