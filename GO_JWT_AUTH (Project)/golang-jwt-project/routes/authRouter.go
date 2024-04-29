package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prudviyedla/golang-jwt-project/controllers"
)

func AuthRoutes(ar *gin.Engine) {

	ar.POST("/users/signup", controller.Signup())
	ar.POST("/users/login", controller.Login())

}
