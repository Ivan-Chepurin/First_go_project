package psql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"main/pkg/models"
)

type SnippetModel struct {
	DB *sqlx.DB
}

func (sm *SnippetModel) Insert(title, content, expires string) (int, error) {
	tx := sm.DB.MustBegin()
	var lastInsertId int
	err := tx.QueryRowx(
		`INSERT INTO snippets (title, content, created, expires)
		VALUES ($1, $2, current_timestamp, current_timestamp + $3)
		RETURNING id`,
		title,
		content,
		fmt.Sprintf("%v day", expires),
	).Scan(&lastInsertId)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (sm *SnippetModel) Get(id int) (models.Snippet, error) {
	var s models.Snippet
	err := sm.DB.Get(&s, `SELECT * FROM snippets WHERE CURRENT_TIMESTAMP < expires AND id=$1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, models.ErrNoRecord
		}
		return s, err
	}
	return s, nil
}

func (sm *SnippetModel) Latest() ([]*models.Snippet, error) {
	var rows []*models.Snippet
	err := sm.DB.Select(
		&rows,
		`SELECT id, title, content, created, expires 
		FROM snippets WHERE  expires > CURRENT_TIMESTAMP 
		ORDER BY created DESC LIMIT 10`,
	)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
