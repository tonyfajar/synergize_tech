// controllers/user_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synergize_tech/repositories"
)

type RequestBodyWallet struct {
	UserId uint    `json:"userId"`
	Amount float64 `json:"amount"`
}

func TopUpWallet(c *gin.Context) {

	var requestBodyWallet RequestBodyWallet
	if err := c.ShouldBindJSON(&requestBodyWallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userExists, _ := repositories.FindById(requestBodyWallet.UserId)
	if userExists == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Exists"})
		return
	}

	repositories.TopUpWallet(requestBodyWallet.UserId, requestBodyWallet.Amount)

	c.JSON(http.StatusCreated, "Success")
}
