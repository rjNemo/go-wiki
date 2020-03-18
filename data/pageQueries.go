package data

const (
	// CreateTablePages is the SQL command used to create the pages table if it
	// not exists.
	CreateTablePages = `
	CREATE TABLE IF NOT EXISTS pages (
	id SERIAL PRIMARY KEY,
	title TEXT UNIQUE NOT NULL,
	body TEXT)
	`

	// GetPage is the SQL command used to retrieve a user from the table
	GetPage = `
	SELECT * FROM pages WHERE title = $1;
	`

	// GetAllPages is the SQL command used to retrieve all pages from the database.
	GetAllPages = `
	SELECT * FROM pages ORDER BY title;
	`

	// FindPages is the SQL command used to retrieve all pages from the table
	// which satisfy the specified condition
	FindPages = `
	SELECT * FROM pages WHERE $1 = $2;
	`

	// InsertPage is the SQL command used to insert a user in the table.
	// Returning the new ID.
	InsertPage = `
	INSERT INTO pages (title, body)
	VALUES ($1, $2)
	RETURNING id
	`

	// UpdatePage is the SQL command used to update a user in the database.
	UpdatePage = `
	UPDATE pages
	SET title = $2, body = $3
	WHERE id = $1;
	`

	// DeletePage is the SQL command used to delete a user from the database.
	DeletePage = `
	DELETE FROM pages
	WHERE id = $1
	RETURNING *;
	`

	// ExistsPage is the SQL command used to check whether a page from the database.
	ExistsPage = `
	SELECT EXISTS(SELECT * FROM pages WHERE title = $1)
	`
)
