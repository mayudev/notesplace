package repository

import (
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
)

func (r *Repository) GetNotebook(id string) (*model.Notebook, error) {
	return r.store.GetNotebook(id)
}

func (r *Repository) CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error {
	return r.store.CreateNotebook(id, protection, hash)
}

func (r *Repository) DeleteNotebook(id string) error {
	return r.store.DeleteNotebook(id)
}
