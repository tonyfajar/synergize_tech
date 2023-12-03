// main.go
package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/synergize_tech/controllers"
	"github.com/synergize_tech/db"
	"github.com/synergize_tech/models"
)

func main() {
	// Create a Postgres database
	db.InitDatabase()

	// Migrate the User table
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Account{})
	db.DB.AutoMigrate(&models.Wallet{})

	// Create a new Gin router
	r := gin.Default()

	r.Use(secure.New(secure.Config{}))

	// Define routes
	r.POST("/login", controllers.LoginUser)
	r.POST("/logout", authMiddleware(), controllers.LogOutUser)
	r.POST("/users", authMiddleware(), controllers.CreateUser)
	r.POST("/account", authMiddleware(), controllers.CreateAccount)
	r.POST("/top_up", authMiddleware(), controllers.TopUpWallet)
	r.GET("/users", authMiddleware(), controllers.GetAllUsers)
	r.GET("/users/:id", authMiddleware(), controllers.GetDetailsUser)
	r.GET("/get_token", controllers.GetToken)

	// Start the server
	r.Run(":8080")

}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = []byte("synergize_tech")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve claims from token"})
			c.Abort()
			return
		}

		// Set the username in the context for later use
		c.Set("username", claims["username"])
	}
}
