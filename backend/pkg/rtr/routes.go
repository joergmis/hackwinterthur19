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
	// authorized := router.Group("/", gin.BasicAuth(users))

	// create tag
	router.POST("/tags", func(c *gin.Context) {
		tag := &db.Tag{}
		c.Bind(&tag)
		tag = db.CreateTag(conn, tag)
		c.JSON(200, structs.Map(tag))
	})

	// search route
	router.POST("/search", func(c *gin.Context) {
		params := c.Request.URL.Query()
		var docs []*db.Document
		if len(params) == 0 {
			c.JSON(512, "no params")
			return
		}
		for _, p := range params {
			for _, m := range p {
				docs = db.SearchForDocuments(conn, m)
				c.JSON(200, docs)
				return
			}
		}
		c.JSON(200, docs)
	})

	// create issue tag
	router.POST("/issuetags", func(c *gin.Context) {
		issueTag := &db.IssueTag{}
		c.Bind(&issueTag)
		issueTag = db.CreateIssueTag(conn, issueTag)
		c.JSON(200, structs.Map(issueTag))
	})

	// create document tag
	router.POST("/documenttags", func(c *gin.Context) {
		documentTag := &db.DocumentTag{}
		c.Bind(&documentTag)
		documentTag = db.CreateDocumentTag(conn, documentTag)
		c.JSON(200, structs.Map(documentTag))
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

	// check if user has correct password
	router.POST("/users/authenticate", func(c *gin.Context) {
		user := &db.User{}
		c.Bind(&user)
		users, err := db.GetUsers(conn)
		if err != nil {
			log.Fatal(err)
		}
		if users[user.Name] == user.Password {
			c.JSON(200, structs.Map(user))
			return
		}
		c.JSON(200, "login not successful")
	})

	// get all issues from the database
	router.GET("/issues", func(c *gin.Context) {
		issues := db.GetAllIssues(conn)
		log.Print(issues)
		// add test issue if database is empty
		if len(issues) == 0 {
			type test struct {
				Name string `json:"name"`
			}
			c.JSON(200, []test{test{Name: "hello"}, test{Name: "hello"}})
		} else {
			c.JSON(200, issues)
		}
	})

	// create an issue
	router.POST("/issues", func(c *gin.Context) {
		issue := &db.Issue{}
		c.Bind(&issue)
		issue = db.InsertIssue(conn, issue)
		c.JSON(200, structs.Map(issue))
	})

	// get a specifig issue
	router.GET("/issues/:id", func(c *gin.Context) {
		issue := db.GetSpecIssue(conn, c.Param("id"))
		c.JSON(200, structs.Map(issue))
	})

	// delete a specific issue
	router.DELETE("/issues/:id", func(c *gin.Context) {
		db.DeleteSpecIssue(conn, c.Param("id"))
		c.JSON(200, gin.H{"delete": "success"})
	})

	// file upload
	router.POST("/fileupload", func(c *gin.Context) {
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
	router.POST("/documents", func(c *gin.Context) {
		document := &db.Document{}
		c.Bind(&document)
		document = db.InsertDocument(conn, document)
		c.JSON(200, structs.Map(document))
	})

	// create a file
	router.POST("/files", func(c *gin.Context) {
		file := &db.File{}
		c.Bind(&file)
		file = db.InsertFile(conn, file)
		c.JSON(200, structs.Map(file))
	})

	// create a note
	router.POST("/notes", func(c *gin.Context) {
		note := &db.Note{}
		c.Bind(&note)
		note = db.InsertNote(conn, note)
		c.JSON(200, structs.Map(note))
	})

	return router
}
