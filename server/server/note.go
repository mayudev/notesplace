package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
	"github.com/mayudev/notesplace/server/validation"
)

// getNoteEndpoint returns a note
func (s *Server) getNoteEndpoint(c *gin.Context) {
	id := c.Param("id")

	note, exists := s.store.GetNote(id)

	if !exists {
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		})
		return
	}

	// Fetch notebook
	notebook, exists := s.store.GetNotebook(note.NotebookID)

	// Check if read access is required
	if notebook.ProtectionLevel > 1 {
		c.JSON(401, util.Response{
			Status:  "error",
			Message: util.Unauthorized,
		})
		return
	}

	// Return note

	c.JSON(200, note)
}

// putNoteEndpoint creates or updates a note
func (s *Server) putNoteEndpoint(c *gin.Context) {
	var body model.Note

	if ok := bindData(c, &body); !ok {
		return
	}

	// Validate the request
	if ok := validation.ValidateNoteUpdate(&body); !ok {
		badRequest(c)
		return
	}

	// Find associated notebook
	notebook, exists := s.store.GetNotebook(body.NotebookID)

	// Check if notebook is protected against writes
	if notebook.ProtectionLevel > 0 {
		// TODO Check if user has write access in the notebook
		unauthorized(c)
		return
	}

	// If ID was not specified, proceed to create a new note.
	if body.ID == "" {
		note := newNote(&body)
		s.store.CreateNote(note)
		c.JSON(201, note)
		return
	}

	// Try to find note in database
	note, exists := s.store.GetNote(body.ID)

	if !exists {
		notFound(c)
		return
	}

	// Verify if notebook ID provided in the request is the same
	// as the actual notebook ID
	if note.NotebookID != body.NotebookID {
		c.JSON(403, util.Response{
			Status:  "error",
			Message: util.Forbidden,
		})
		return
	}

	// Set last update time
	body.UpdatedAt = time.Now()

	// Update the note
	result, err := s.store.UpdateNote(body)

	if err != nil {
		internalServerError(c)
	}

	c.JSON(200, result)

	// TODO validation.ValidateNoteUpdate(body)

	/* if body.ID == "" {
		// Create a new note
		// TODO validation

		// TODO
		note := model.Note{
			ID:         "", // todo
			NotebookID: body.NotebookID,
			Title:      body.Title,
			Order:      body.Order, // TODO calculate order
			Content:    body.Content,
			CreatedAt:  time.UnixMicro(0),
			UpdatedAt:  time.UnixMicro(0),
		}

		err := s.store.CreateNote(note)

		if err != nil {
			c.JSON(500, util.Response{
				Status:  "error",
				Message: "error", // todo
			})
			return
		}

		c.JSON(201, note)
		return
	} */

	// TODO validate contents
	// TODO DRY
	// TODO dont ignore error
}

func (s *Server) deleteNoteEndpoint(c *gin.Context) {
	id := c.Param("id")

	// TODO DRY
	note, exists := s.store.GetNote(id)

	if !exists {
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NoteNotFound,
		})
		return
	}

	// Fetch notebook
	notebook, exists := s.store.GetNotebook(note.NotebookID)
	if !exists {
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		})
		return
	}

	// Check if write access is required
	if notebook.ProtectionLevel > 0 {
		c.JSON(401, util.Response{
			Status:  "error",
			Message: util.Unauthorized,
		})
		return
	}

	s.store.DeleteNote(id)
	c.String(200, "Ok")
	return
}

func newNote(note *model.Note) *model.Note {
	// TODO validation

	// TODO
	mew := model.Note{
		ID:         util.GenerateID().String(), // todo
		NotebookID: note.NotebookID,
		Title:      note.Title,
		Order:      nil, // TODO calculate order
		Content:    note.Content,
		CreatedAt:  time.UnixMicro(0),
		UpdatedAt:  time.UnixMicro(0),
	}

	return &mew
}
