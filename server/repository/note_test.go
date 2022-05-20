package repository_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/repository"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
)

func TestGetNote(t *testing.T) {
	store := test.StubServerStore{
		Notes: map[string]model.Note{
			"note": {
				ID: "note",
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("find an existing note", func(t *testing.T) {
		got, exists := repo.GetNote("note")

		assert.Equal(t, got.ID, "note")
		assert.True(t, exists)
	})

	t.Run("returns false on a note that does not exist", func(t *testing.T) {
		_, exists := repo.GetNote("other")

		assert.False(t, exists)
	})
}

func TestUpdateNote(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"notebook": {
				ID:   "notebook",
				Name: "notebook",
			},
		},
		Notes: map[string]model.Note{
			"note_1": {
				ID:         "note_1",
				NotebookID: "notebook",
				Title:      "Note 1",
				Order:      0,
			},
			"note_2": {
				ID:         "note_2",
				NotebookID: "notebook",
				Title:      "Note 2",
				Order:      1,
			},
			"note_3": {
				ID:         "note_3",
				NotebookID: "notebook",
				Title:      "Note 3",
				Order:      2,
			},
			"note_4": {
				ID:         "note_4",
				NotebookID: "notebook",
				Title:      "Note 3",
				Order:      3,
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("updates a note's content", func(t *testing.T) {
		note1 := store.Notes["note_1"]
		note1.Content = "Note 1 content"

		repo.UpdateNote(note1)

		assert.Equal(t, "Note 1 content", store.Notes["note_1"].Content)
	})

	t.Run("updates a note's title", func(t *testing.T) {
		note1 := store.Notes["note_1"]
		note1.Title = "Note 1 title"

		repo.UpdateNote(note1)

		assert.Equal(t, "Note 1 title", store.Notes["note_1"].Title)
	})

	t.Run("reorders notes when necessary", func(t *testing.T) {
		note1 := store.Notes["note_1"]
		note1.Order = 2

		repo.UpdateNote(note1)

		assert.Equal(t, uint(2), store.Notes["note_1"].Order)
		assert.Equal(t, uint(1), store.Notes["note_3"].Order)
		assert.Equal(t, uint(0), store.Notes["note_2"].Order)
	})

	t.Run("reorders notes when necessary (down)", func(t *testing.T) {
		note3 := store.Notes["note_3"]
		note3.Order = 0

		repo.UpdateNote(note3)

		assert.Equal(t, uint(0), store.Notes["note_3"].Order)
		assert.Equal(t, uint(1), store.Notes["note_2"].Order)
	})
}

func TestCreateNote(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"notebook": {
				ID:   "notebook",
				Name: "notebook",
			},
		},
		Notes: map[string]model.Note{
			"note_1": {
				ID:         "note_1",
				NotebookID: "notebook",
				Title:      "Note 1",
				Order:      0,
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("creates a new note and assigns correct order to it", func(t *testing.T) {
		newNote := &model.Note{
			ID:         "note_2", // TODO insert ID generation in repository
			NotebookID: "notebook",
			Title:      "Note 2",
		}

		err := repo.CreateNote(newNote)
		assert.NoError(t, err)

		assert.Equal(t, "note_2", store.Notes["note_2"].ID)
		assert.Equal(t, uint(1), store.Notes["note_2"].Order)
	})

}

func TestDeleteNote(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"notebook": {
				ID:   "notebook",
				Name: "notebook",
			},
		},
		Notes: map[string]model.Note{
			"note_1": {
				ID:         "note_1",
				NotebookID: "notebook",
				Title:      "Note 1",
				Order:      0,
			},
			"note_2": {
				ID:         "note_2",
				NotebookID: "notebook",
				Title:      "Note 2",
				Order:      1,
			},
			"note_3": {
				ID:         "note_3",
				NotebookID: "notebook",
				Title:      "Note 3",
				Order:      2,
			},
			"note_4": {
				ID:         "note_4",
				NotebookID: "notebook",
				Title:      "Note 3",
				Order:      3,
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("deletes a note and reorders other notes accordingly", func(t *testing.T) {
		note2, exists := repo.GetNote("note_2")
		assert.True(t, exists)

		err := repo.DeleteNote(&note2)
		assert.NoError(t, err)

		assert.Equal(t, uint(0), store.Notes["note_1"].Order)
		assert.Equal(t, uint(1), store.Notes["note_3"].Order)
		assert.Equal(t, uint(2), store.Notes["note_4"].Order)
	})
}
