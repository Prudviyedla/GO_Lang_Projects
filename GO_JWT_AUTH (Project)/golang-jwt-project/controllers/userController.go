package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prudviyedla/golang-jwt-project/database"
	helper "github.com/prudviyedla/golang-jwt-project/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection("user")
var validate = validator.New()

func Hashpassword() {
	// Implement Hashpassword function
}

func VerifyPassword() {
	// Implement VerifyPassword function
}

func Signup() {
	// Implement Signup function
}

func Login() {
	// Implement Login function
}

func GetUsers() {
	// Implement GetUsers function
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		// Example usage of MatchUserTypeToUid function
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Your logic for retrieving user information based on userId
		// This function should query your database (e.g., MongoDB) to get user details
		// Then return the user details in the response
	}
}
