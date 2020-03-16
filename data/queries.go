package data

const (
	// QueryCreateTable is the SQL command used to create the users table if it
	// not exists.
	QueryCreateTable = `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	age INT,
	first_name TEXT,
	last_name TEXT,
	email TEXT UNIQUE NOT NULL)
	`
	// QueryInsertUser is the SQL command used to insert a user in the table.
	// Returning the new ID.
	QueryInsertUser = `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
)
