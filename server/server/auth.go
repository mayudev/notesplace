package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/auth"
)

func (s *Server) authenticateEndpoint(c *gin.Context) {
	id := c.GetHeader("Notebook")
	password := c.GetHeader("Password")

	if len(id) == 0 || len(password) == 0 {
		badRequest(c)
		return
	}

	notebook, exists := s.store.GetNotebook(id)

	if !exists {
		notFound(c)
		return
	}

	match := auth.ComparePassword(notebook.Password, password)

	if match {
		c.String(200, "")
	} else {
		c.String(401, "")
	}
}
