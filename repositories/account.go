package repositories

import (
	"github.com/synergize_tech/db"
	"github.com/synergize_tech/models"
)

func CreateAccount(account *models.Account) error {

	return db.DB.Create(account).Error

}

func CheckAccount(account *models.Account) (*models.Account, error) {

	result := db.DB.Where("account = ? AND bank = ? AND user_id = ? ", account.Account, account.Bank, account.UserId).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func GetAllAccount(UserId uint) ([]models.Account, error) {
	var account []models.Account
	db.DB.Where("user_id = ? ", UserId).Find(&account)
	return account, nil
}
