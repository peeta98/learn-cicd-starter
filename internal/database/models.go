// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"time"
)

type Note struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Note      string
	UserID    string
}

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}