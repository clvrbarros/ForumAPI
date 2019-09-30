package models

import "database/sql"

// Conn type is a struct for connections
type Conn struct {
	db *sql.DB
}

// Connect to the database
func Connect(db *sql.DB) *Conn {
	return &Conn{db}
}