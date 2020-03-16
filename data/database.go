package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rjNemo/go-wiki/settings"
)

// Connect read the connection parameters to establish a connection to the
// database.
func Connect() {
	connStr := settings.ConnStr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")
}
