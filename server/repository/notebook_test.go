package repository_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/repository"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
)

func TestGetNotebook(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"noteboook": {
				ID: "noteboook",
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("find an existing notebook", func(t *testing.T) {
		got, err := repo.GetNotebook("noteboook")

		assert.Equal(t, got.ID, "noteboook")
		assert.NoError(t, err)
	})

	t.Run("returns false on a notebook that does not exist", func(t *testing.T) {
		_, err := repo.GetNotebook("other")

		assert.Error(t, err)
	})
}

func TestCreateNotebook(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"noteboook": {
				ID: "noteboook",
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("creates a new notebook", func(t *testing.T) {
		err := repo.CreateNotebook("new", "test", 1, "")
		assert.NoError(t, err)

		assert.Equal(t, "new", store.Notebooks["new"].ID)
		assert.Equal(t, auth.ProtectionLevel(1), store.Notebooks["new"].ProtectionLevel)
		assert.Equal(t, "", store.Notebooks["new"].Password)
	})
}

func TestDeleteNotebook(t *testing.T) {
	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"noteboook": {
				ID: "noteboook",
			},
		},
	}

	repo := repository.NewRepository(&store)

	t.Run("deletes a notebook", func(t *testing.T) {
		err := repo.DeleteNotebook("notebook")
		assert.NoError(t, err)

		_, err = repo.GetNotebook("notebook")
		assert.Error(t, err)
	})
}
