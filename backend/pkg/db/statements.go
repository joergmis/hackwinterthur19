// Package db provides the functions and constants to select / insert
// values into the database
// notes for the future:
package db

import (
	"fmt"
	"log"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

const (
	// create table statements
	createUserTable        string = `create table if not exists user(id integer primary key, name text, password text);`
	createIssueTable       string = `create table if not exists issue(id integer primary key, name text, description text, userid int, fileid int, documentid int, foreign key(userid) references user(id) on delete cascade, foreign key(fileid) references file(id) on delete cascade, foreign key(documentid) references document(id) on delete cascade);`
	createFileTable        string = `create table if not exists file(id integer primary key, location text, documentid int, foreign key(documentid) references document(id) on delete cascade);`
	createNoteTable        string = `create table if not exists note(id integer primary key, content text, fileid int, foreign key(fileid) references file(id) on delete cascade);`
	createDocumentTable    string = `create table if not exists document(id integer primary key, name text, text text, location text);`
	createTagTable         string = `create table if not exists tag(id integer primary key, name text);`
	createIssueTagTable    string = `create table if not exists issuetag(id integer primary key, issueid int, tagid int, foreign key(issueid) references issue(id) on delete cascade, foreign key(tagid) references tag(id) on delete cascade);`
	createDocumentTagTable string = `create table if not exists documenttag(id integer primary key, documentid int, tagid int, foreign key(documentid) references document(id) on delete cascade, foreign key(tagid) references tag(id) on delete cascade);`
	// insert statements
	insertUser        string = `insert into user(name, password) values (?,?);`
	insertFile        string = `insert into file(location, documentid) values (?,?);`
	insertFileWithout string = `insert into file(location) values (?);`
	insertNote        string = `insert into note(content, fileid) values (?,?);`
	insertIssue       string = `insert into issue(name, description, userid, fileid, documentid) values (?,?,?,?,?);`
	insertDocument    string = `insert into document(name, text, location) values (?,?,?);`
	insertTag         string = `insert into tag(name) values (?);`
	insertIssueTag    string = `insert into issuetag(issueid, tagid) values (?,?);`
	insertDocumentTag string = `insert into documenttag(documentid, tagid), values (?,?);`
	// select statements
	selectUsers         string = `select user.name, user.password from user;`
	selectIssues        string = `select * from issue;`
	selectSpecificIssue string = `select * from issue where issue.id = ?;`
	selectFileLocation  string = `select file.location from file where file.id = ?;`
	// delete statements
	deleteIssue string = `delete from issue where issue.id = ?;`
)

// CreateTag creates an issue
func CreateTag(conn *sqlite3.Conn, tag *Tag) *Tag {
	err := conn.Exec(insertTag, tag.Name)
	if err != nil {
		log.Fatal(err)
	}
	tag.ID = int(conn.LastInsertRowID())
	return tag
}

// CreateIssueTag creates an issue
func CreateIssueTag(conn *sqlite3.Conn, issueTag *IssueTag) *IssueTag {
	err := conn.Exec(insertIssueTag, issueTag.Issueid, issueTag.Tagid)
	if err != nil {
		log.Fatal(err)
	}
	issueTag.ID = int(conn.LastInsertRowID())
	return issueTag
}

// CreateDocumentTag creates an issue
func CreateDocumentTag(conn *sqlite3.Conn, documentTag *DocumentTag) *DocumentTag {
	err := conn.Exec(insertDocumentTag, documentTag.Documentid, documentTag.Tagid)
	if err != nil {
		log.Fatal(err)
	}
	documentTag.ID = int(conn.LastInsertRowID())
	return documentTag
}

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
	if err != nil {
		return err
	}
	err = conn.Exec(createIssueTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createTagTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createIssueTagTable)
	if err != nil {
		return err
	}
	err = conn.Exec(createDocumentTagTable)
	return err
}

// InsertIssue inserts an issue
func InsertIssue(conn *sqlite3.Conn, issue *Issue) *Issue {
	log.Print(issue)
	err := conn.Exec(insertIssue, issue.Name, issue.Description, issue.Userid, issue.Fileid, issue.Documentid)
	if err != nil {
		log.Fatal(err)
	}
	issue.ID = int(conn.LastInsertRowID())
	return issue
}

// InsertUser inserts a user
func InsertUser(conn *sqlite3.Conn, user *User) *User {
	err := conn.Exec(insertUser, user.Name, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = int(conn.LastInsertRowID())
	return user
}

// DeleteSpecIssue delets an issue
func DeleteSpecIssue(conn *sqlite3.Conn, id string) {
	stmt, err := conn.Prepare(deleteIssue, id)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// InsertDocument creates a new document
func InsertDocument(conn *sqlite3.Conn, document *Document) *Document {
	err := conn.Exec(insertDocument, document.Name, document.Text, document.Location)
	if err != nil {
		log.Fatal(err)
	}
	document.ID = int(conn.LastInsertRowID())
	return document
}

// InsertFile creates a new file
func InsertFile(conn *sqlite3.Conn, file *File) *File {
	log.Print(file.Location)
	err := conn.Exec(insertFileWithout, file.Location)
	if err != nil {
		log.Fatal(err)
	}
	file.ID = int(conn.LastInsertRowID())
	return file
}

// InsertNote creates a new note
func InsertNote(conn *sqlite3.Conn, note *Note) *Note {
	err := conn.Exec(insertNote, note.Content, note.Fileid)
	if err != nil {
		log.Fatal(err)
	}
	note.ID = int(conn.LastInsertRowID())
	return note
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
	users["user"] = "password"
	return users, err
}

// GetSpecIssue returns a specific issue
func GetSpecIssue(conn *sqlite3.Conn, id string) *Issue {
	issue := &Issue{}
	stmt, err := conn.Prepare(selectSpecificIssue, id)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for {
		_, err := stmt.Step()
		if err != nil {
			log.Fatal(err)
		}

		err = stmt.Scan(&issue.ID, &issue.Name, &issue.Description, &issue.Userid, &issue.Fileid, &issue.Documentid)
		log.Println(issue.Name)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
	return issue
}

// GetAllIssues returns all issues in the database
func GetAllIssues(conn *sqlite3.Conn) []Issue {
	issues := []Issue{}
	stmt, err := conn.Prepare(selectIssues)
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
		var id int
		var name string
		var description string
		var userid int
		var fileid int
		var documentid int
		err = stmt.Scan(&id, &name, &description, &userid, &fileid, &documentid)
		if err != nil {
			log.Fatal(err)
		}
		issue := &Issue{ID: id, Name: name, Description: description, Userid: userid, Fileid: fileid, Documentid: documentid}
		issues = append(issues, *issue)
	}
	_ = stmt.Reset()
	return issues
}
