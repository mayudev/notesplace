package repository

import (
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
)

type Store interface {
	// Notebook
	GetNotebook(id string) (*model.Notebook, error)
	CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error
	DeleteNotebook(id string) error

	// Note
	GetNote(id string) (*model.Note, error)
	CreateNote(payload *model.Note) error
	UpdateNote(payload *model.Note) (model.Note, error)
	DeleteNote(id string) error

	// Internal
	GetNoteByOrder(notebook string, order uint) (*model.Note, error)
	NoteCount(notebook string) uint
}
