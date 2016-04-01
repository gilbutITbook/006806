package models

import (
	"time"
)

type Comment struct {
	Id        int64     `db:"id"`
	Body      string    `db:"body"`
	Commenter string    `db:"commenter"`
	PostId    int       `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
