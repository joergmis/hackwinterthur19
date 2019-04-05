package rtr

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"../db"
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// InitRouter initialises the routes
func InitRouter(users map[string]string, conn *sqlite3.Conn) *gin.Engine {
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
		user = db.InsertUser(conn, user)
		log.Printf("user got id '%d'", user.ID)
		c.JSON(200, structs.Map(user))
	})

	// get all issues from the database
	authorized.GET("/issues", func(c *gin.Context) {
		issues := db.GetAllIssues(conn)
		log.Print(issues)
		c.JSON(200, structs.Map(issues))
	})

	// create an issue
	authorized.POST("/issues", func(c *gin.Context) {
		issue := &db.Issue{}
		c.Bind(&issue)
		issue = db.InsertIssue(conn, issue)
		c.JSON(200, structs.Map(issue))
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
