package main

import (
	"fmt"
	"log"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func main() {
	conn, err := sqlite3.Open("./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("hello world")
}
