package db

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

var (
	test *sqlite3.Stmt
)

const (
	createUserTable     string = `create table if not exists user(id int primary key, name text, password text);`
	createFileTable     string = `create table if not exists file(id int primary key, location text, documentid int, foreign key(documentid) references document(id));`
	createNoteTable     string = `create table if not exists note(id int primary key, content text, fileid int, foreign key(fileid) references file(id));`
	createDocumentTable string = `create table if not exists document(id int primary key, name text, location text);`
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

// CreateStmts prepares all the statements
func CreateStmts(conn *sqlite3.Conn) {
	// stmt, err := conn.Prepare()
}
