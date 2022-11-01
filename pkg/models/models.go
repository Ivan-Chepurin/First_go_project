package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

var SnipSchema = `
	CREATE TABLE IF NOT EXISTS snippets (
    id INTEGER,
    title VARCHAR(100),
    content TEXT,
    created TIMESTAMP WITH TIME ZONE,
    expires TIMESTAMP WITH TIME ZONE 
)`

type Snippet struct {
	Id      int       `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Created time.Time `db:"created"`
	Expires time.Time `db:"expires"`
}
