package repository

import (
	"fmt"

	"github.com/mayudev/notesplace/server/model"
)

func (r *Repository) GetNote(id string) (model.Note, bool) {
	return r.store.GetNote(id)
}

func (r *Repository) UpdateNote(data model.Note) (model.Note, error) {
	note, ok := r.store.GetNote(data.ID)

	if !ok {
		return model.Note{}, fmt.Errorf("note not found")
	}

	if data.Content != "" {
		note.Content = data.Content
	}

	if data.Title != "" {
		note.Title = data.Title
	}

	if data.Order != note.Order {
		if data.Order > note.Order {
			// Note has been moved __up__
			for i := note.Order + 1; i <= data.Order; i++ {
				existing, exists := r.store.GetNoteByOrder(data.NotebookID, i)
				if !exists {
					return model.Note{}, fmt.Errorf("note not found")
				}
				existing.Order--
				r.store.UpdateNote(existing)
			}
		} else {
			// Note has been moved __down__
			for i := note.Order - 1; i >= data.Order; i-- {
				existing, exists := r.store.GetNoteByOrder(data.NotebookID, i)
				if !exists { // corrupted notebook uh oh
					return model.Note{}, fmt.Errorf("note not found")
				}
				existing.Order++
				r.store.UpdateNote(existing)
			}
		}
		note.Order = data.Order
	}

	return r.store.UpdateNote(note)
}

func (r *Repository) CreateNote(data *model.Note) error {
	return r.store.CreateNote(data)
}

func (r *Repository) DeleteNote(data *model.Note) error {
	count := r.store.NoteCount(data.NotebookID)

	for i := data.Order; i < count; i++ {
		note, exists := r.store.GetNoteByOrder(data.NotebookID, i)
		if !exists {
			continue
		}
		note.Order--
		r.store.UpdateNote(note)
	}

	return r.store.DeleteNote(data.ID)
}
