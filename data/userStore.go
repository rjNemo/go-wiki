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

// GetAll retrieves all the users from the database.
func (us UserStore) GetAll() ([]model.User, error) {
	var id, age int
	var firstName, lastName, email string

	rows, err := us.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		err := rows.Scan(&id, &age, &firstName, &lastName, &email)
		if err != nil {
			log.Fatal(err) // too severe
		}
		u := model.NewUser(id, age, firstName, lastName, email)
		users = append(users, u)
	}
	return users, nil
}

// Find retrieves a user from the database using an expression
func (us UserStore) Find(k interface{}, v interface{}) ([]model.User, error) {
	var id, age int
	var firstName, lastName, email string

	rows, err := us.db.Query(QueryFind, k, v)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		err := rows.Scan(&id, &age, &firstName, &lastName, &email)
		if err != nil {
			log.Fatal(err) // too severe
		}
		u := model.NewUser(id, age, firstName, lastName, email)
		users = append(users, u)
	}
	return users, nil
}

// Get retrieves the user identified by 'id' from database
func (us UserStore) Get(uid int) (model.User, error) {
	var id, age int
	var firstName, lastName, email string

	row := us.db.QueryRow(QueryGet, uid)
	switch err := row.Scan(&id, &age, &firstName, &lastName, &email); err {
	case sql.ErrNoRows:
		log.Println("No entry returned")
		return model.User{}, err
	case nil:
		return model.NewUser(id, age, firstName, lastName, email), nil
	default:
		log.Fatal(err)
		return model.User{}, err
	}
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
		log.Fatal("Delete failed") // too severe
	}
}
