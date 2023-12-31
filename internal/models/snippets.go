package models

import (
	"database/sql"
	"time"
)

// define a Snippet type to hold the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define SnippetModel type to wrap a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Return specific snippet based on its ID
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return nil, nil
}

// Return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
