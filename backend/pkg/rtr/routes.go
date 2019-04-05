package rtr

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// InitRouter initialises the routes
func InitRouter(users map[string]string) *gin.Engine {
	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(users))

	// dummy route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// create user (unauthorized!)
	router.POST("/users", func(c *gin.Context) {
		user := &user{}
		c.Bind(&user)
		log.Print(user)
		c.JSON(200, gin.H{
			"success": "true",
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
