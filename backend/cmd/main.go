package main

import (
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

	// get user gets either the users or adds
	// a default user:
	// user: user, password: password
	users, err := db.GetUsers(conn)
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := rtr.InitRouter(users)
	router.Run(":8090")
}
