package data

import (
	"database/sql"
	"log"
	"strings"

	"github.com/rjNemo/go-wiki/models"
)

// PageStore is used to perform page-related CRUD operations on the DB
type PageStore struct {
	db *sql.DB
}

// NewPageStore constructs a PageStore connected to the database.
func NewPageStore(db *sql.DB) PageStore {
	return PageStore{db: db}
}

// CreateTable creates the page table in the database if it exits not yet.
func (ps PageStore) CreateTable() {
	if _, err := ps.db.Exec(CreateTablePages); err != nil {
		log.Fatal(err)
	}
	// log.Print("Table successfully created!")
}

// GetAll retrieves all the pages from the database.
func (ps PageStore) GetAll() ([]models.Page, error) {
	var id int
	var title string
	var body []byte

	rows, err := ps.db.Query(GetAllPages)
	if err != nil {
		return nil, err
	}

	var pages []models.Page
	for rows.Next() {
		err := rows.Scan(&id, &title, &body)
		if err != nil {
			log.Fatal(err) // too severe
		}
		p := models.NewPage(id, title, body)
		pages = append(pages, *p)
	}
	return pages, nil
}

// Get retrieves the page identified by 'id' from database
func (ps PageStore) Get(t string) (models.Page, error) {
	var id int
	var title string
	var body []byte
	t = strings.Title(t)
	row := ps.db.QueryRow(GetPage, t)
	switch err := row.Scan(&id, &title, &body); err {
	case sql.ErrNoRows:
		log.Println("No entry returned")
		return models.Page{}, err
	case nil:
		return *models.NewPage(id, title, body), nil
	default:
		log.Fatal(err)
		return models.Page{}, err
	}
}

// Add inserts a page in the database.
func (ps PageStore) Add(u models.Page) {
	id := 0
	err := ps.db.QueryRow(InsertPage, u.Title(), u.Body()).Scan(&id)
	if err != nil {
		log.Fatal(err) // too severe
	}
	log.Println("New Page ID is:", id)
}

// Update edits page identified by 'id' in the database
func (ps PageStore) Update(id int, u models.Page) {
	res, err := ps.db.Exec(UpdatePage, id, u.Title(), string(u.Body()))
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

// Delete removes page identified by 'id' from the database
func (ps PageStore) Delete(id int) {
	res, err := ps.db.Exec(QueryDelete, id)
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

// Exists returns true if the page named title exists, false otherwise
func (ps PageStore) Exists(title string) (e bool) {
	err := ps.db.QueryRow(ExistsPage, strings.Title(title)).Scan(&e)
	if err != nil {
		log.Fatal(err) // too severe
	}
	return e
}

// Find retrieves a page from the database using an expression
// func (ps PageStore) Find(k interface{}, v interface{}) ([]models.Page, error) {
// 	var id, title int
// 	var body, lastName, email string

// 	rows, err := ps.db.Query(QueryFind, k, v)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var pages []models.Page
// 	for rows.Next() {
// 		err := rows.Scan(&id, &title, &body, &lastName, &email)
// 		if err != nil {
// 			log.Fatal(err) // too severe
// 		}
// 		u := models.NewPage(id, title, body, lastName, email)
// 		pages = append(pages, u)
// 	}
// 	return pages, nil
// }
