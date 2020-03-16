package data

import (
	"database/sql"
	"log"

	"github.com/rjNemo/go-wiki/model"
)

// UserStore is used to perform user-related CRUD operations on the DB
type UserStore struct {
	db *sql.DB
}

// NewUserStore constructs a UserStore connected to the database.
func NewUserStore(db *sql.DB) UserStore {
	return UserStore{db: db}
}

// Add inserts a user in the database.
func (us UserStore) Add(u model.User) {
	id := 0
	err := us.db.QueryRow(QueryInsert, u.Age(), u.Email(), u.FirstName(), u.LastName()).Scan(&id)
	if err != nil {
		log.Fatal(err) // too severe
	}
	log.Println("New User ID is:", id)
}

// Update edits user identified by 'id' in the database
func (us UserStore) Update(id int, u model.User) {
	res, err := us.db.Exec(QueryUpdate, id, u.Age(), u.Email(), u.FirstName(), u.LastName())
	if err != nil {
		log.Fatal(err) // too severe
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err) // too severe
	}
	if count == 0 {
		log.Fatal("Update failed") // too severe
	}
}


// Delete removes user identified by 'id' from the database
func (us UserStore) Delete(id int) {
	res, err := us.db.Exec(QueryDelete, id)
	if err != nil {
		log.Fatal(err) // too severe
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err) // too severe
	}
	if count == 0 {
		log.Fatal("Update failed") // too severe
	}