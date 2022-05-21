package database_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/database"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
)

var DB *database.Database

func TestMain(m *testing.M) {
	DB = database.NewDatabase("postgres://postgres:secret@localhost:5432/testing")

	m.Run()

	DB.Cleanup()
	DB.Close()
}

func TestNotebook(t *testing.T) {
	notebookID := "testing_notebook1"

	t.Run("creates a notebook", func(t *testing.T) {
		err := DB.CreateNotebook(notebookID, auth.ProtectionLevel(0), "")
		assert.NoError(t, err)
	})

	t.Run("returns previously created notebook", func(t *testing.T) {
		got, exists := DB.GetNotebook(notebookID)
		assert.True(t, exists)

		want := model.Notebook{
			ID:              notebookID,
			Name:            "TODO", // TODO name handling kek
			Password:        "",
			ProtectionLevel: 0,
			CreatedAt:       got.CreatedAt,
			UpdatedAt:       got.UpdatedAt,
		}

		test.AssertDeepEqual(t, got, want)
	})

	t.Run("deletes that notebook", func(t *testing.T) {
		err := DB.DeleteNotebook(notebookID)
		assert.NoError(t, err)
	})

	t.Run("returns an error when a notebook doesn't exist", func(t *testing.T) {
		_, exists := DB.GetNotebook(notebookID)
		assert.False(t, exists)
	})

	t.Run("returns an error when trying to delete a notebook that doesn't exist", func(t *testing.T) {
		err := DB.DeleteNotebook(notebookID)
		assert.Error(t, err)
	})

}