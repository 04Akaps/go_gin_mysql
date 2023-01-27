// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Diary struct {
	ID        int64          `json:"id"`
	Content   sql.NullString `json:"content"`
	UserEmail string         `json:"user_email"`
	CreatedAt time.Time      `json:"created_at"`
}

type User struct {
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Age       int64     `json:"age"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
