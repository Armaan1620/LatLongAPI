package models

import "time"

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never serialize password
	CreatedAt time.Time `json:"created_at"`
}

// UserStore defines the interface for user storage operations
type UserStore interface {
	CreateUser(email, hashedPassword string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
}

