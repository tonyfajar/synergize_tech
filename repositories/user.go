package repositories

import (
	"github.com/synergize_tech/db"
	"github.com/synergize_tech/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return db.DB.Create(user).Error

}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := db.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindById(userId uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var user []models.User
	if err := db.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(username string, password string) (*models.User, error) {
	var user models.User

	result := db.DB.Where("username = ? ", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
