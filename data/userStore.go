package data

import (
	"database/sql"
	"log"

	"github.com/rjNemo/go-wiki/models"
)

// UserStore is used to perform user-related CRUD operations on the DB
type UserStore struct {
	db *sql.DB
}

// NewUserStore constructs a UserStore connected to the database.
func NewUserStore(db *sql.DB) UserStore {
	return UserStore{db: db}
}

// CreateTable creates the user table in the database if it exits not yet.
func (us UserStore) CreateTable() {
	if _, err := us.db.Exec(QueryCreateTable); err != nil {
		log.Fatal(err)
	}
	// log.Print("Table successfully created!")
}

// GetAll retrieves all the users from the database.
func (us UserStore) GetAll() ([]models.User, error) {
	var id, age int
	var firstName, lastName, email string

	rows, err := us.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		err := rows.Scan(&id, &age, &firstName, &lastName, &email)
		if err != nil {
			log.Fatal(err) // too severe
		}
		u := models.NewUser(id, age, firstName, lastName, email)
		users = append(users, u)
	}
	return users, nil
}

// Find retrieves a user from the database using an expression
// func (us UserStore) Find(k interface{}, v interface{}) ([]models.User, error) {
// 	var id, age int
// 	var firstName, lastName, email string

// 	rows, err := us.db.Query(QueryFind, k, v)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []models.User
// 	for rows.Next() {
// 		err := rows.Scan(&id, &age, &firstName, &lastName, &email)
// 		if err != nil {
// 			log.Fatal(err) // too severe
// 		}
// 		u := models.NewUser(id, age, firstName, lastName, email)
// 		users = append(users, u)
// 	}
// 	return users, nil
// }

// Get retrieves the user identified by 'id' from database
func (us UserStore) Get(uid int) (models.User, error) {
	var id, age int
	var firstName, lastName, email string

	row := us.db.QueryRow(QueryGet, uid)
	switch err := row.Scan(&id, &age, &firstName, &lastName, &email); err {
	case sql.ErrNoRows:
		log.Println("No entry returned")
		return models.User{}, err
	case nil:
		return models.NewUser(id, age, firstName, lastName, email), nil
	default:
		log.Fatal(err)
		return models.User{}, err
	}
}

// Add inserts a user in the database.
func (us UserStore) Add(u models.User) {
	id := 0
	err := us.db.QueryRow(QueryInsert, u.Age(), u.Email(), u.FirstName(), u.LastName()).Scan(&id)
	if err != nil {
		log.Fatal(err) // too severe
	}
	log.Println("New User ID is:", id)
}

// Update edits user identified by 'id' in the database
func (us UserStore) Update(id int, u models.User) {
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
