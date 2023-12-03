// controllers/user_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synergize_tech/models"
	"github.com/synergize_tech/repositories"
)

type RequestBodyAccount struct {
	UserId  uint   `json:"userId"`
	Account string `json:"account"`
	Bank    string `json:"bank"`
	Name    string `json:"name"`
}

func CreateAccount(c *gin.Context) {

	var requestBodyAccount RequestBodyAccount
	if err := c.ShouldBindJSON(&requestBodyAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{
		UserId:  requestBodyAccount.UserId,
		Account: requestBodyAccount.Account,
		Bank:    requestBodyAccount.Bank,
		Name:    requestBodyAccount.Name,
	}

	userExists, _ := repositories.FindById(requestBodyAccount.UserId)
	if userExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Exists"})
		return
	}
	_, err := repositories.CheckAccount(&account)
	if err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account Already Exists"})
		return
	}

	repositories.CreateAccount(&account)

	c.JSON(http.StatusCreated, account)
}
