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
	IssueID int `json:"IssueID"`
	TagID   int `json:"TagID"`
}

// DocumentTag representation
type DocumentTag struct {
	ID         int `json:"ID"`
	DocumentID int `json:"DocumentID"`
	TagID      int `json:"TagID"`
}

// Issue representation
type Issue struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	UserID      int    `json:"UserID"`
	FileID      int    `json:"FileID"`
	DocumentID  int    `json:"DocumentID"`
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
	DocumentID int    `json:"DocumentID"`
}

// Note representation
type Note struct {
	ID      int    `json:"ID"`
	Content string `json:"Content"`
	FileID  int    `json:"FileID"`
}
