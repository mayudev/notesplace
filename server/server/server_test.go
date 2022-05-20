package server_test

import (
	"time"

	"github.com/mayudev/notesplace/server/auth"
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

func (store *StubServerStore) CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error {
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

func (store *StubServerStore) GetNoteByOrder(notebook string, order uint) (model.Note, bool) {
	for _, v := range store.notes {
		if v.Order == order {
			return v, true
		}
	}

	return model.Note{}, false
}

func (store *StubServerStore) UpdateNote(data model.Note) (model.Note, error) {
	store.notes[data.ID] = data

	return store.notes[data.ID], nil
}

func (store *StubServerStore) CreateNote(data *model.Note) error {
	store.notes[data.ID] = *data

	return nil
}

func (store *StubServerStore) DeleteNote(id string) error {
	delete(store.notes, id)
	return nil
}
