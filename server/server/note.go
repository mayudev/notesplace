package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
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

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, util.Response{
			Status:  "error",
			Message: util.RequestInvalid,
		})
		return
	}

	// Fetch notebook
	notebook, exists := s.store.GetNotebook(body.NotebookID)

	// Check if write access is required
	if notebook.ProtectionLevel > 0 {
		c.JSON(401, util.Response{
			Status:  "error",
			Message: util.Unauthorized,
		})
		return
		// Authentication required
	}

	// TODO validation.ValidateNoteUpdate(body)

	note, exists := s.store.GetNote(body.ID)

	if !exists {
		// TODO Create a note
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		})
		return
	}

	// Check if note comes from a different notebook
	if note.NotebookID != body.NotebookID {
		c.JSON(403, util.Response{
			Status:  "error",
			Message: util.Forbidden,
		})
	}

	// TODO validate contents
	// TODO DRY
	// TODO dont ignore error
	result, _ := s.store.UpdateNote(body)

	c.JSON(200, result)
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
