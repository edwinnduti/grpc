package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Dob	   string
	Password  string
	Description string
	CreatedAt string
	UpdatedAt string
}

type UserCreated struct {
	ID        string
}