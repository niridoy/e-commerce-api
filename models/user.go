package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles []Role `json:"roles"` // Added Roles to User
}
