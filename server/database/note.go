package database

import (
	"context"
	"fmt"
	"time"

	"github.com/mayudev/notesplace/server/model"
)

func (d *Database) GetNote(id string) (*model.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	note := new(model.Note)

	err := d.conn.QueryRow(ctx, "SELECT id, notebook_id, title, \"order\", content, created_at, updated_at FROM notesplace.notes WHERE id = $1", id).Scan(&note.ID, &note.NotebookID, &note.Title, &note.Order, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (d *Database) CreateNote(data *model.Note) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO notesplace.notes
(id, notebook_id, title, "order", content, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := d.conn.Exec(ctx, query, data.ID, data.NotebookID, data.Title, data.Order, data.Content, time.Now().UTC(), time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateNote(data *model.Note) (model.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE notesplace.notes
SET title = $2, "order" = $3, content = $4, updated_at = $5
WHERE id = $1`

	_, err := d.conn.Exec(ctx, query, data.ID, data.Title, data.Order, data.Content, time.Now().UTC())
	if err != nil {
		return model.Note{}, err
	}

	// TODO
	return model.Note{}, err
}

func (d *Database) DeleteNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := d.conn.Exec(ctx, "DELETE FROM notesplace.notes WHERE id = $1", id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	return nil
}

func (d *Database) GetNoteByOrder(notebook string, order uint) (model.Note, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	note := new(model.Note)

	err := d.conn.QueryRow(ctx, "SELECT id, notebook_id, title, \"order\", content, created_at, updated_at FROM notesplace.notes WHERE \"order\" = $1 AND notebook_id = $2", order, notebook).Scan(&note.ID, &note.NotebookID, &note.Title, &note.Order, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return model.Note{}, false
	}

	return *note, true
}

func (d *Database) NoteCount(notebook string) uint {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	num := new(uint)
	err := d.conn.QueryRow(ctx, "SELECT COUNT(*) FROM notesplace.notes WHERE notebook_id = $1", notebook).Scan(&num)

	if err != nil {
		return 0
	}

	return *num
}
