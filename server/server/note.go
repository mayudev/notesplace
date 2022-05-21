package server

import (
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
		notFound(c)
		return
	}

	// Fetch notebook
	notebook, exists := s.store.GetNotebook(note.NotebookID)

	if !exists {
		notFound(c)
		return
	}

	// Check if read access is required
	if notebook.ProtectionLevel.Protected() {
		valid := s.Validate(c, notebook.ID)

		if !valid {
			unauthorized(c)
			return
		}
	}

	// Return note
	c.JSON(200, note)
}

// putNoteEndpoint creates or updates a note
func (s *Server) putNoteEndpoint(c *gin.Context) {
	var body model.Note

	if ok := bindData(c, &body); !ok {
		badRequest(c)
		return
	}

	// Validate the request
	if ok := validation.ValidateNoteUpdate(&body); !ok {
		badRequest(c)
		return
	}

	// Find associated notebook
	notebook, exists := s.store.GetNotebook(body.NotebookID)

	if !exists {
		notFound(c)
		return
	}

	// Check if notebook is protected against writes
	if notebook.ProtectionLevel.WriteProtected() {
		valid := s.Validate(c, notebook.ID)

		if !valid {
			unauthorized(c)
			return
		}
	}

	// If ID was not specified, proceed to create a new note.
	if body.ID == "" {
		note, err := s.store.CreateNote(&body)

		if err != nil {
			internalServerError(c)
			return
		}

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

	// Update the note
	result, err := s.store.UpdateNote(body)

	if err != nil {
		internalServerError(c)
		return
	}

	c.JSON(200, result)
}

func (s *Server) deleteNoteEndpoint(c *gin.Context) {
	id := c.Param("id")

	// TODO DRY
	note, exists := s.store.GetNote(id)

	if !exists {
		notFound(c)
		return
	}

	// Fetch notebook
	notebook, exists := s.store.GetNotebook(note.NotebookID)
	if !exists {
		notFound(c)
		return
	}

	// Check if write access is required
	if notebook.ProtectionLevel.WriteProtected() {
		valid := s.Validate(c, notebook.ID)

		if !valid {
			unauthorized(c)
			return
		}
	}

	s.store.DeleteNote(&note)
	c.String(200, "Ok")
	return
}
