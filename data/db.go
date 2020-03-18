package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresql database package
)

// NewDB read the connection parameters to establish a connection to the
// database.
func NewDB(connection string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connection to database successfully established!")
	return db, nil
}

// Context registers the application data stores
type Context struct {
	Pages PageStore
	Users UserStore
}

// NewContext returns a context instance
func NewContext(db *sql.DB) Context {
	pageStore := NewPageStore(db)
	userStore := NewUserStore(db)
	return Context{Pages: pageStore, Users: userStore}
}

// Store interface defines the methods any store must satisfy
type Store interface {
	CreateTable()
	Add(i interface{})
	Get(i interface{}) (interface{}, error)
	GetAll() (interface{}, error)
	Delete(id int)
	// Find(ex string)
	Update(id int, i interface{})
}
