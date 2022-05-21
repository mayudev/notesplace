package database

import (
	"context"
	"fmt"
	"time"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
)

func (d *Database) GetNotebook(id string) (model.Notebook, bool) {
	notebook := new(model.Notebook)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := d.conn.QueryRow(ctx, "SELECT id, title, password, protection_level, created_at, updated_at FROM notesplace.notebooks WHERE id = $1", id).Scan(&notebook.ID, &notebook.Name, &notebook.Password, &notebook.ProtectionLevel, &notebook.CreatedAt, &notebook.UpdatedAt)
	if err != nil {
		return model.Notebook{}, false
	}

	return *notebook, true
}

// TODO proper error handling, title
func (d *Database) CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO notesplace.notebooks
	(id, title, password, protection_level, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := d.conn.Exec(ctx, query, id, "TODO", "", 0, time.Now().UTC(), time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteNotebook(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := d.conn.Exec(ctx, "DELETE FROM notesplace.notebooks WHERE id = $1", id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	return nil
}
