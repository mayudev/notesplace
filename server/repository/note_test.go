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
		got, err := repo.GetNote("note")

		assert.Equal(t, got.ID, "note")
		assert.NoError(t, err)
	})

	t.Run("returns an error on a note that does not exist", func(t *testing.T) {
		_, got := repo.GetNote("other")

		assert.Error(t, got)
	})
}

func TestGetNotesByNotebook(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"noteboook": {
				ID: "noteboook",
			},
		},
		Notes: map[string]model.Note{
			"note1": {
				ID:         "note1",
				NotebookID: "notebook",
			},
			"note2": {
				ID:         "note2",
				NotebookID: "notebook",
			},
			"note3": {
				ID:         "note3",
				NotebookID: "notebook",
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("returns all notes in a notebook", func(t *testing.T) {
		got, err := repo.GetNotesByNotebook("notebook")
		assert.NoError(t, err)

		assert.Len(t, got, 3)

		assert.Equal(t, "note1", got[0].ID)
		assert.Equal(t, "note2", got[1].ID)
		assert.Equal(t, "note3", got[2].ID)
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

		update := model.Note{
			ID:         note1.ID,
			NotebookID: note1.NotebookID,
			Order:      2,
		}

		repo.UpdateNote(update)

		assert.Equal(t, uint(2), store.Notes["note_1"].Order)
		assert.Equal(t, uint(1), store.Notes["note_3"].Order)
		assert.Equal(t, uint(0), store.Notes["note_2"].Order)
	})

	t.Run("reorders notes when necessary (down)", func(t *testing.T) {
		note3 := store.Notes["note_3"]
		update := model.Note{
			ID:         note3.ID,
			NotebookID: note3.NotebookID,
			Order:      0,
		}

		repo.UpdateNote(update)

		assert.Equal(t, uint(0), store.Notes["note_3"].Order)
		assert.Equal(t, uint(1), store.Notes["note_2"].Order)
	})

	t.Run("does not reorder a note if content was specified", func(t *testing.T) {
		note2 := store.Notes["note_2"]
		update := model.Note{
			ID:         note2.ID,
			NotebookID: note2.NotebookID,
			Content:    "a",
			Title:      "a",
			Order:      0,
		}

		repo.UpdateNote(update)

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
			NotebookID: "notebook",
			Title:      "Note 2",
		}

		note, err := repo.CreateNote(newNote)
		assert.NoError(t, err)

		assert.Equal(t, "notebook", store.Notes[note.ID].NotebookID)
		assert.Equal(t, uint(1), store.Notes[note.ID].Order)
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
		note2, err := repo.GetNote("note_2")
		assert.NoError(t, err)

		err = repo.DeleteNote(note2)
		assert.NoError(t, err)

		assert.Equal(t, uint(0), store.Notes["note_1"].Order)
		assert.Equal(t, uint(1), store.Notes["note_3"].Order)
		assert.Equal(t, uint(2), store.Notes["note_4"].Order)
	})
}
