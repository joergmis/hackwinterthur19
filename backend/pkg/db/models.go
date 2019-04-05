package db

// User representation
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
	Content  string `json:"content"`
	Fileid   int    `json:"fileid"`
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
