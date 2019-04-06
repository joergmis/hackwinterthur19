package db

// User representation
type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

// Tag representation
type Tag struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

// IssueTag representation
type IssueTag struct {
	ID      int `json:"ID"`
	Issueid int `json:"Issueid"`
	Tagid   int `json:"Tagid"`
}

// DocumentTag representation
type DocumentTag struct {
	ID         int `json:"ID"`
	Documentid int `json:"Documentid"`
	Tagid      int `json:"Tagid"`
}

// Issue representation
type Issue struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Userid      int    `json:"Userid"`
	Fileid      int    `json:"Fileid"`
	Documentid  int    `json:"Documentid"`
}

// Document representation
type Document struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Text     string `json:"Text"`
	Location string `json:"Location"`
}

// File representation
type File struct {
	ID         int    `json:"ID"`
	Location   string `json:"Location"`
	Documentid int    `json:"Documentid"`
}

// Note representation
type Note struct {
	ID      int    `json:"ID"`
	Content string `json:"Content"`
	Fileid  int    `json:"Fileid"`
}
