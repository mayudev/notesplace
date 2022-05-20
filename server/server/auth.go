package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/util"
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
		token, err := s.issuer.Issue(id)
		if err != nil {
			internalServerError(c)
			return
		}

		c.String(200, token)
	} else {
		c.JSON(401, util.Response{
			Status:  "error",
			Message: util.IncorrectPassword,
		})
	}
}
