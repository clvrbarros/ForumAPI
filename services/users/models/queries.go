package models

// Users queries
const (
	queryUserInsert = `
		INSERT INTO users (
		email,
		password,
		first_name,
		last_name,
		country,
		active
		)
		VALUES($1, $2, $3, $4, $5, false)
	`
)

// Auth queires
const (
	queryAuthPassword = "SELECT password FROM users WHERE email=$1"
)