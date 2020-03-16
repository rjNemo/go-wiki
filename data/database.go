package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresql database package
	"github.com/rjNemo/go-wiki/model"
	"github.com/rjNemo/go-wiki/settings"
)

// Connect read the connection parameters to establish a connection to the
// database.
func Connect() {
	log.Print("Inside connect func")
	connStr := settings.ConnStr
	db, err := sql.Open("postgres", connStr)
	log.Print("opened connect")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Print("try to ping")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")

	createTable(db)
	u := model.TestUser()
	insertUser(db, u)
	log.Print(u)
}

func sqlExec(db *sql.DB, s string) {
	if _, err := db.Exec(s); err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Command successfully executed!: %s", s)
}

func createTable(db *sql.DB) {
	sqlExec(db, QueryCreateTable)
	log.Print("Table successfully created!")
}

func insertUser(db *sql.DB, u model.User) {
	id := 0
	err := db.QueryRow(QueryInsertUser, u.Age(), u.Email(), u.FirstName(), u.LastName()).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	// u.SetID(id)
	log.Println("New User ID is:", id)
}

func insertUsers(db *sql.DB) {
	query := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, 32, "ruidy.nemausat@gmail.com", "Ruidy", "Nemausat")
	if err != nil {
		log.Fatal(err)
	}
}
