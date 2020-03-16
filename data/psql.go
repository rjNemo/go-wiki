package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresql database package
	"github.com/rjNemo/go-wiki/model"
	"github.com/rjNemo/go-wiki/settings"
)

// UsePSQL read the connection parameters to establish a connection to the
// database.
func UsePSQL() {
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

	createUserTable(db)
	store := NewUserStore(db)
	// u := model.TestUser()
	// store.Add(u)
	// log.Print(u)
	u1 := model.NewUser(3, 19, "paul", "newman", "PdsN@FDKML.COM")
	store.Update(16, u1)
}

func sqlExec(db *sql.DB, s string) {
	if _, err := db.Exec(s); err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Command successfully executed!: %s", s)
}

func createUserTable(db *sql.DB) {
	sqlExec(db, QueryCreateTable)
	log.Print("Table successfully created!")
}

// Store interface defines the methods any store must satisfy
type Store interface {
	CreateTable()
	Add(i interface{})
	Get()
	Find(id int)
	Update(id int, i interface{})
	Delete(id int)
}
