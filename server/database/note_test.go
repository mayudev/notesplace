package database_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/test"
	"github.com/mayudev/notesplace/server/util"
	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	testing1 := "testing_1"
	testing2 := "testing_2"

	validNote1 := util.GenerateID().String()
	validNote2 := util.GenerateID().String()

	assert.NotEqual(t, validNote1, validNote2)

	DB.CreateNotebook(testing1, auth.ProtectionLevel(0), "")
	DB.CreateNotebook(testing2, auth.ProtectionLevel(0), "")

	t.Run("returns valid note count when there are no notes", func(t *testing.T) {
		count := DB.NoteCount(testing1)
		assert.Equal(t, uint(0), count)
	})

	t.Run("creates a note", func(t *testing.T) {
		note := model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "Notebook title 1",
			Order:      0,
			Content:    "Notebook content 1",
		}

		err := DB.CreateNote(&note)
		assert.NoError(t, err)
	})

	t.Run("find a note", func(t *testing.T) {
		got, err := DB.GetNote(validNote1)
		assert.NoError(t, err)

		want := &model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "Notebook title 1",
			Order:      0,
			Content:    "Notebook content 1",
			CreatedAt:  got.CreatedAt,
			UpdatedAt:  got.UpdatedAt,
		}

		test.AssertDeepEqual(t, got, want)
	})

	t.Run("find the note by order", func(t *testing.T) {
		got, err := DB.GetNoteByOrder(testing1, 0)
		assert.NoError(t, err)

		want := &model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "Notebook title 1",
			Order:      0,
			Content:    "Notebook content 1",
			CreatedAt:  got.CreatedAt,
			UpdatedAt:  got.UpdatedAt,
		}

		test.AssertDeepEqual(t, got, want)
	})

	t.Run("returns valid note count with two notes", func(t *testing.T) {
		note := model.Note{
			ID:         validNote2,
			NotebookID: testing1,
			Title:      "Notebook title 2",
			Order:      1,
			Content:    "Notebook content 2",
		}

		err := DB.CreateNote(&note)
		assert.NoError(t, err)

		count := DB.NoteCount(testing1)
		assert.Equal(t, uint(2), count)
	})

	t.Run("does not find a note that doesn't exist", func(t *testing.T) {
		_, err := DB.GetNote("whatever")
		assert.Error(t, err)
	})

	t.Run("updates a note", func(t *testing.T) {
		update := model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "New title 1",
			Order:      1,
			Content:    "New content 1",
		}

		_, err := DB.UpdateNote(&update)
		assert.NoError(t, err)

		got, err := DB.GetNote(validNote1)
		assert.NoError(t, err)

		want := &model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "New title 1",
			Order:      1,
			Content:    "New content 1",
			CreatedAt:  got.CreatedAt,
			UpdatedAt:  got.UpdatedAt,
		}

		test.AssertDeepEqual(t, got, want)
	})

	t.Run("deletes a note", func(t *testing.T) {
		err := DB.DeleteNote(validNote1)
		assert.NoError(t, err)

		_, err = DB.GetNote(validNote1)
		assert.Error(t, err)
	})
}
