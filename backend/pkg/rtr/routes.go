package rtr

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

// InitRouter initialises the routes
func InitRouter(users map[string]string) *gin.Engine {
	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(users))

	// add routes
	authorized.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// file upload
	authorized.POST("/fileupload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create(path.Join(".", "tmp", filename))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})

	return router
}
