package rtr

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"../db"
	"github.com/gin-gonic/gin"
)

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
		user := &db.User{}
		c.Bind(&user)
		log.Printf("user '%s' registered with password '%s'", user.Name, user.Password)
		// TODO: insert user into database
		c.JSON(200, gin.H{
			"success": "true",
		})
	})

	// get all issues from the database
	authorized.GET("/issues", func(c *gin.Context) {
		// TODO:
		c.JSON(200, gin.H{"hello": "world", "bye": "world"})
	})

	// create an issue
	authorized.POST("/issues", func(c *gin.Context) {
		// TODO:

		c.JSON(200, gin.H{"hello": "world", "bye": "world"})
	})

	// get a specifig issue
	authorized.GET("/issues/id", func(c *gin.Context) {
		// TODO:
		c.JSON(200, gin.H{"hello": "world", "bye": "world"})
	})

	// delete a specific issue
	authorized.DELETE("/issues/id", func(c *gin.Context) {
		// TODO:
		c.JSON(200, gin.H{"hello": "world", "bye": "world"})
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
