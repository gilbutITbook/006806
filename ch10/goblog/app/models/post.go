package models

import (
	"time"
)

type Post struct {
	Id        int       `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Comments  []Comment
}
