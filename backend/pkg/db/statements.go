// Package db provides the functions and constants to select / insert
// values into the database
// notes for the future:
// LastInsertRowID() returns int64 of the last successful insert statement
package db

import (
	"fmt"
	"log"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

const (
	// create table statements
	createUserTable     string = `create table if not exists user(id int primary key, name text, password text);`
	createFileTable     string = `create table if not exists file(id int primary key, location text, documentid int, foreign key(documentid) references document(id));`
	createNoteTable     string = `create table if not exists note(id int primary key, content text, fileid int, foreign key(fileid) references file(id));`
	createDocumentTable string = `create table if not exists document(id int primary key, name text, location text);`
	// insert statements
	insertUser     string = `insert into user(name, password) values (?,?);`
	insertFile     string = `insert into file(location, documentid) values (?,?);`
	insertNote     string = `insert into note(content, fileid) values (?,?);`
	insertDocument string = `insert into document(name, location) values (?,?);`
	// select statements
	selectUsers        string = `select user.name, user.password from user;`
	selectFileLocation string = `select file.location from file where file.id = ?;`
)

// CreateTables creates the tables if they don't exist
func CreateTables(conn *sqlite3.Conn) error {
	err := conn.Exec(createUserTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createDocumentTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createFileTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createNoteTable)
	return err
}

// GetUsers returns a map with all the users which
// are in the database
func GetUsers(conn *sqlite3.Conn) (map[string]string, error) {
	users := make(map[string]string)
	stmt, err := conn.Prepare(selectUsers)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for {
		hasRow, err := stmt.Step()
		if err != nil {
			log.Fatal(err)
		}
		if !hasRow {
			// The query is finished
			break
		}

		// Use Scan to access column data from a row
		var name string
		var password string
		err = stmt.Scan(&name, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s, password: %s\n", name, password)
		users[name] = password
	}
	if len(users) == 0 {
		users = map[string]string{"user": "password"}
	}
	return users, err
}
