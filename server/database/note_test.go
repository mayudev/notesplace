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
		got, exists := DB.GetNote(validNote1)
		assert.True(t, exists)

		want := model.Note{
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

	t.Run("does not find a note that doesn't exist", func(t *testing.T) {
		_, exists := DB.GetNote("whatever")
		assert.False(t, exists)
	})

	t.Run("updates a note", func(t *testing.T) {
		update := model.Note{
			ID:         validNote1,
			NotebookID: testing1,
			Title:      "New title 1",
			Order:      1,
			Content:    "New content 1",
		}

		_, err := DB.UpdateNote(update)
		assert.NoError(t, err)

		got, exists := DB.GetNote(validNote1)
		assert.True(t, exists)

		want := model.Note{
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

		_, exists := DB.GetNote(validNote1)
		assert.False(t, exists)
	})
}
