package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
}

func NewDatabase(url string) *Database {
	conn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to db %v\n", err)
		os.Exit(1)
	}

	err = runMigrations(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to migrate tables %v\n", err)
		os.Exit(1)
	}

	return &Database{conn: conn}
}

func (d *Database) Close() {
	d.conn.Close()
}

// Cleanup cleans up the database. Used only for testing.
func (d *Database) Cleanup() {
	d.conn.Exec(context.Background(), "DROP SCHEMA notesplace CASCADE;")
}

var notebooksMigration = `
CREATE TABLE IF NOT EXISTS notesplace.notebooks (
	id VARCHAR(32) PRIMARY KEY,
	title VARCHAR(256),
	password TEXT,
	protection_level SMALLINT NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
	);
`

var notesMigration = `
CREATE TABLE IF NOT EXISTS notesplace.notes (
	id uuid PRIMARY KEY,
	notebook_id VARCHAR(32) REFERENCES notesplace.notebooks(id) ON DELETE CASCADE,
	title VARCHAR(256),
	"order" INT NOT NULL,
	content TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
	);`

func runMigrations(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), "CREATE SCHEMA IF NOT EXISTS notesplace;")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), notebooksMigration)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), notesMigration)
	if err != nil {
		return err
	}

	return nil
}
