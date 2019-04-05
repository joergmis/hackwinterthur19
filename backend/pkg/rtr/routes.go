package rtr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initialises the routes
func InitRouter(users map[string]string) *gin.Engine {
	router := gin.Default()

	authorized := router.Group("/", gin.BasicAuth(users))

	// add routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})

	return router
}
