package model

import "strings"

// User represent a go-wiki user. It encapsulate its unique identifier, first and
// last names, email and age
type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	age       int
}

// TestUser constructs a sample John Doe user. For Dev and Test only
func TestUser() User {
	return User{
		id:        42,
		firstName: "John",
		lastName:  "Doe",
		email:     "jddddSZ@mail.com",
		age:       42,
	}
}

// NewUser is a constructor
func NewUser(id, age int, first, last, email string) User {
	return User{
		id:        id,
		firstName: strings.Title(first),
		lastName:  strings.Title(last),
		email:     strings.ToLower(email),
		age:       age,
	}
}

// ID is a getter
func (u User) ID() int {
	return u.id
}

// FirstName is a getter
func (u User) FirstName() string {
	return u.firstName
}

// LastName is a getter
func (u User) LastName() string {
	return u.lastName
}

// Email is a getter
func (u User) Email() string {
	return u.email
}

// SetEmail is a setter
func (u *User) SetEmail(email string) {
	u.email = strings.ToLower(email)
}

// Age is a getter
func (u User) Age() int {
	return u.age
}
