// controllers/user_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/synergize_tech/db"
	"github.com/synergize_tech/models"
	"github.com/synergize_tech/repositories"
)

type RequestBodyUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestBodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestBodyLogout struct {
	Username string `json:"username"`
}

type RequestBodyGetDetails struct {
	UserId uint `form:"id"`
}

type ResponseDetailsUser struct {
	Name     string            `json:"name"`
	Username string            `json:"username"`
	Balance  float64           `json:"balance"`
	Account  []ResponseAccount `json:"account"`
}

type ResponseLogin struct {
	Description string `json:"description"`
	Username    string `json:"username"`
	Token       string `json:"token"`
}

type ResponseAccount struct {
	Name    string `json:"name"`
	Bank    string `json:"bank"`
	Account string `json:"account"`
}

func CreateUser(c *gin.Context) {

	var requestBodyUser RequestBodyUser
	if err := c.ShouldBindJSON(&requestBodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exists, _ := repositories.FindByUsername(requestBodyUser.Username)
	if exists != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Already Exists"})
		return
	}
	user := models.User{
		Username: requestBodyUser.Username,
		Password: requestBodyUser.Password,
		Name:     requestBodyUser.Name,
	}

	repositories.CreateUser(&user)

	wallet := models.Wallet{
		UserId:  user.ID,
		Balance: 0,
	}
	repositories.CreateWallet(&wallet)

	c.JSON(http.StatusCreated, user)
}

func GetToken(c *gin.Context) {
	username := "root"

	token, err := createToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func createToken(username string) (string, error) {
	var secretKey = []byte("synergize_tech")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetAllUsers(c *gin.Context) {
	users, _ := repositories.GetAllUsers()

	c.JSON(http.StatusCreated, users)
}

func GetDetailsUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, _ := repositories.FindById(uint(id))
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Exists"})
		return
	}

	wallet, _ := repositories.GetDetailsWallet(uint(id))
	account, _ := repositories.GetAllAccount(uint(id))

	var accountArray []ResponseAccount
	for _, acc := range account {
		accObj := ResponseAccount{
			Account: acc.Account,
			Bank:    acc.Bank,
			Name:    acc.Name,
		}
		accountArray = append(accountArray, accObj)
	}

	response := ResponseDetailsUser{
		Name:     user.Name,
		Username: user.Username,
		Balance:  wallet.Balance,
		Account:  accountArray,
	}
	c.JSON(http.StatusOK, response)
}

func LoginUser(c *gin.Context) {
	db.InitRedis()

	var requestBodyLogin RequestBodyLogin
	if err := c.ShouldBindJSON(&requestBodyLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := repositories.LoginUser(requestBodyLogin.Username, requestBodyLogin.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Exists or Password didnt match"})
		return
	}
	token, err := createToken(user.Username)

	responseLogin := ResponseLogin{
		Description: "Success",
		Username:    user.Username,
		Token:       token,
	}

	errRedis := db.SetKey(user.Username, token, 3600)
	if errRedis != nil {
		fmt.Println("Error setting key:", err)
		return
	}

	c.JSON(http.StatusCreated, responseLogin)
}

func LogOutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
