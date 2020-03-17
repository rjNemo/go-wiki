package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresql database package
	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/settings"
)

// UsePSQL read the connection parameters to establish a connection to the
// database.
func UsePSQL() *sql.DB {
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

	log.Println("Connection to database successfully established!")

	// store := NewUserStore(db)
	// store.CreateTable()

	return db
	// u := models.TestUser()
	// store.Add(u)
	// log.Print(u)
	// u1 := models.NewUser(3, 20, "paul", "newman", "PdsNz@FDKML.COM")
	// store.Update(16, u1)
	// log.Println(store.Get(1))
	// store.Delete(8)
	// log.Println(store.GetAll())
	// log.Println(store.Find("first_name", "John"))
}

// Store interface defines the methods any store must satisfy
type Store interface {
	CreateTable()
	Add(i interface{})
	Get(id int) (models.User, error)
	GetAll(id int) ([]models.User, error)
	Delete(id int)
	// Find(ex string)
	Update(id int, i interface{})
}
