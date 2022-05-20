package repository

import (
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
)

type Store interface {
	// Notebook
	GetNotebook(id string) (model.Notebook, bool)
	CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error
	DeleteNotebook(id string) error

	// Note
	GetNote(id string) (model.Note, bool)
	CreateNote(data *model.Note) error
	UpdateNote(data model.Note) (model.Note, error)
	DeleteNote(id string) error

	// Internal
	GetNoteByOrder(notebook string, order uint) (model.Note, bool)
}
