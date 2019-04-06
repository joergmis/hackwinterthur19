package rtr

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"../db"
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// InitRouter initialises the routes
func InitRouter(users map[string]string, conn *sqlite3.Conn) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Datatype",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
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
	authorized.GET("/issues/:id", func(c *gin.Context) {
		issue := db.GetSpecIssue(conn, c.Param("id"))
		c.JSON(200, structs.Map(issue))
	})

	// delete a specific issue
	authorized.DELETE("/issues/:id", func(c *gin.Context) {
		db.DeleteSpecIssue(conn, c.Param("id"))
		c.JSON(200, gin.H{"delete": "success"})
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

	// create a document
	authorized.POST("/documents", func(c *gin.Context) {
		document := &db.Document{}
		c.Bind(&document)
		document = db.InsertDocument(conn, document)
		c.JSON(200, structs.Map(document))
	})

	// create a file
	authorized.POST("/files", func(c *gin.Context) {
		file := &db.File{}
		c.Bind(&file)
		file = db.InsertFile(conn, file)
		c.JSON(200, structs.Map(file))
	})

	// create a note
	authorized.POST("/notes", func(c *gin.Context) {
		note := &db.Note{}
		c.Bind(&note)
		note = db.InsertNote(conn, note)
		c.JSON(200, structs.Map(note))
	})

	return router
}
