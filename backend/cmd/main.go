package main

import (
	"fmt"
	"log"

	"../pkg/db"
	"../pkg/rtr"
	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func main() {
	conn, err := sqlite3.Open("./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create tables
	err = db.CreateTables(conn)
	if err != nil {
		log.Fatal(err)
	}

	router := rtr.InitRouter()
	router.Run(":8080")

	fmt.Println("hello world")
}
