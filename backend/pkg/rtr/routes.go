package rtr

import "github.com/gin-gonic/gin"

// InitRouter initialises the routes
func InitRouter() *gin.Engine {
	router := gin.Default()

	// add routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
