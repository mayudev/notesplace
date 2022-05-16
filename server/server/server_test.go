package server_test

import (
	"time"

	"github.com/mayudev/notesplace/server/model"
)

type StubServerStore struct {
	notebooks map[string]model.Notebook
}

func NewStubServerStore() *StubServerStore {
	store := StubServerStore{}
	store.notebooks = map[string]model.Notebook{}

	return &store
}

func (store *StubServerStore) GetNotebook(id string) (model.Notebook, bool) {
	value, ok := store.notebooks[id]

	if !ok {
		return model.Notebook{}, false
	}

	return value, true
}

func (store *StubServerStore) CreateNotebook(id string, protection uint8, hash string) error {
	store.notebooks[id] = model.Notebook{
		ID:              id,
		Name:            id,
		Password:        hash,
		ProtectionLevel: protection,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return nil
}

func (store *StubServerStore) DeleteNotebook(id string) error {
	delete(store.notebooks, id)

	return nil
}
