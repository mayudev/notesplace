package server

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
	"github.com/mayudev/notesplace/server/validation"
)

func (s *Server) getNotebookEndpoint(c *gin.Context) {
	id := c.Param("id")
	notebook, err := s.store.GetNotebook(id)

	if err != nil {
		notFound(c)
		return
	}

	notebook.Password = ""

	// Check if read access is required
	if notebook.ProtectionLevel.Protected() {
		valid := s.Validate(c, id)

		if !valid {
			unauthorized(c)
			return
		}
	}

	// Insert all notes
	notes, err := s.store.GetNotesByNotebook(id)
	if err != nil {
		internalServerError(c)
		return
	}

	notebook.Notes = notes

	c.JSON(200, notebook)
}

func (s *Server) createNotebookEndpoint(c *gin.Context) {
	var req model.NotebookCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c)
		return
	}

	err := validation.ValidateNotebookCreate(req)

	if err != nil {
		badRequest(c)
		return
	}

	password := ""

	// Password is to be set
	if req.ProtectionLevel.WriteProtected() && req.Password != "" {
		hash, err := s.hasher.HashPassword(req.Password)
		if err != nil {
			internalServerError(c)
			return
		}

		password = hash
	}

	id, err := gonanoid.New(12)
	if err != nil {
		internalServerError(c)
		return
	}

	s.store.CreateNotebook(id, "test", req.ProtectionLevel, password)

	c.JSON(200, model.NotebookCreateResponse{
		ID: id,
		Response: util.Response{
			Status:  "success",
			Message: util.NotebookCreated,
		},
	})
	return
}

func (s *Server) deleteNotebookEndpoint(c *gin.Context) {
	id := c.Param("id")
	notebook, err := s.store.GetNotebook(id)

	if err != nil {
		notFound(c)
		return
	}

	// Check if write access is required
	if notebook.ProtectionLevel.WriteProtected() {
		valid := s.Validate(c, id)

		if !valid {
			unauthorized(c)
			return
		}

	}

	// Deleting all notes will be handled by the database
	s.store.DeleteNotebook(id)
	c.String(200, "Ok")
	return
}
