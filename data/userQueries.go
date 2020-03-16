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

	// QueryGet is the SQL command used to retrieve a user from the table
	QueryGet = `
	SELECT * id FROM users WHERE id=$1;
	`

	// QueryInsert is the SQL command used to insert a user in the table.
	// Returning the new ID.
	QueryInsert = `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	// QueryUpdate is the SQL command used to update a user in the database.
	QueryUpdate = `
	UPDATE users
	SET age = $2, email = $3, first_name = $4, last_name = $5
	WHERE id = $1;
	`

	// QueryDelete is the SQL command used to delete a user from the database.
	QueryDelete = `
	DELETE FROM users
	WHERE id = $1;
	`
)
