package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/prudviyedla/golang-jwt-project/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT") //we were retrieving the port value as get environmental from .env file
	if port == "" {
		port = "8000" // if it's doesn't exist there we were taking 8000 port to host
	}

	router := gin.New() //middleware not attached and no handler methods
	//routers:=gin.Default()	//middleware attached and consists all handler methods
	router.Use(gin.Logger()) //Use attaches a global middleware to the router. i. e. the middleware attached through Use() will be included in the handlers chain for every single request. Even 404, 405, static files... For example, this is the right place for a logger or error management middleware.

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)

}
