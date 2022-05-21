package repository

import (
	"fmt"
	"time"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
)

func (r *Repository) GetNote(id string) (*model.Note, error) {
	return r.store.GetNote(id)
}

func (r *Repository) UpdateNote(data model.Note) (model.Note, error) {
	note, err := r.store.GetNote(data.ID)

	if err != nil {
		return model.Note{}, err
	}

	if data.Content != "" {
		note.Content = data.Content
	}

	if data.Title != "" {
		note.Title = data.Title
	}

	// Update Order only if Content and Title were not updated.
	if (data.Content == "" && data.Title == "") && data.Order != note.Order {
		count := r.store.NoteCount(data.NotebookID)
		if data.Order >= count {
			return model.Note{}, fmt.Errorf("incorrect order")
		}

		if data.Order > note.Order {
			// Note has been moved __up__
			for i := note.Order + 1; i <= data.Order; i++ {
				existing, err := r.store.GetNoteByOrder(data.NotebookID, i)
				if err != nil {
					return model.Note{}, fmt.Errorf("note not found")
				}
				existing.Order--
				r.store.UpdateNote(existing)
			}
		} else {
			// Note has been moved __down__
			for i := note.Order - 1; i >= data.Order; i-- {
				existing, err := r.store.GetNoteByOrder(data.NotebookID, i)
				if err != nil { // corrupted notebook uh oh
					return model.Note{}, fmt.Errorf("note not found")
				}
				existing.Order++
				r.store.UpdateNote(existing)

				// prevent overflow
				if i == 0 {
					break
				}
			}
		}
		note.Order = data.Order
	}

	return r.store.UpdateNote(note)
}

func (r *Repository) CreateNote(data *model.Note) (*model.Note, error) {
	count := r.store.NoteCount(data.NotebookID)

	if count > 2000 {
		return &model.Note{}, fmt.Errorf("exceeded note limit")
	}

	// Generate an ID
	id := util.GenerateID().String()
	note := &model.Note{
		ID:         id,
		NotebookID: data.NotebookID,
		Title:      data.Title,
		Order:      count,
		Content:    data.Content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return note, r.store.CreateNote(note)
}

func (r *Repository) DeleteNote(data *model.Note) error {
	count := r.store.NoteCount(data.NotebookID)

	for i := data.Order; i < count; i++ {
		note, err := r.store.GetNoteByOrder(data.NotebookID, i)
		if err != nil {
			continue
		}
		note.Order--
		r.store.UpdateNote(note)
	}

	return r.store.DeleteNote(data.ID)
}
