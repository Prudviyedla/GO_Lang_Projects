package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prudviyedla/golang-jwt-project/controllers"
	"github.com/prudviyedla/golang-jwt-project/middleware"
)

func UserRoutes(ur *gin.Engine) {

	ur.Use(middleware.Authenticate())
	ur.GET("/users", controller.GetUsers())
	ur.GET("/users/:user_id", controller.GetUser())
}
