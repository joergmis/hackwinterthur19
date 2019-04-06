package db

// User representation
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Tag representation
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IssueTag representation
type IssueTag struct {
	ID      int `json:"id"`
	Issueid int `json:"issueid"`
	Tagid   int `json:"tagid"`
}

// DocumentTag representation
type DocumentTag struct {
	ID         int `json:"id"`
	Documentid int `json:"documentid"`
	Tagid      int `json:"tagid"`
}

// Issue representation
type Issue struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Userid      int    `json:"userid"`
	Fileid      int    `json:"fileid"`
	Documentid  int    `json:"documentid"`
}

// Document representation
type Document struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Text     string `json:"text"`
	Location string `json:"location"`
}

// File representation
type File struct {
	ID         int    `json:"id"`
	Location   string `json:"location"`
	Documentid int    `json:"documentid"`
}

// Note representation
type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Fileid  int    `json:"fileid"`
}
