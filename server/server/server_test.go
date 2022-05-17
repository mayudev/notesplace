package server_test

import (
	"fmt"
	"time"

	"github.com/mayudev/notesplace/server/model"
)

type StubServerStore struct {
	notebooks map[string]model.Notebook
	notes     map[string]model.Note
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

func (store *StubServerStore) GetNote(id string) (model.Note, bool) {
	value, ok := store.notes[id]

	if !ok {
		return model.Note{}, false
	}

	return value, true
}

func (store *StubServerStore) UpdateNote(data model.Note) (model.Note, error) {
	note, ok := store.GetNote(data.ID)
	if !ok {
		return model.Note{}, fmt.Errorf("note not found")
	}

	if data.Content != "" {
		note.Content = data.Content
	}

	if data.Title != "" {
		note.Title = data.Title
	}

	if data.Order != nil {
		// TODO reorder other notes (possibly not here)
		note.Order = data.Order
	}

	store.notes[data.ID] = note

	return store.notes[data.ID], nil
}
