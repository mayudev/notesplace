package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
)

func (s *Server) getNotebookEndpoint(c *gin.Context) {
	id := c.Param("id")
	notebook, exists := s.store.GetNotebook(id)

	if !exists {
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		})
		return
	}

	c.JSON(200, notebook)
}

func (s *Server) createNotebookEndpoint(c *gin.Context) {
	var createRequest model.NotebookCreate

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(400, util.Response{
			Status:  "error",
			Message: util.RequestInvalid,
		})
	}

	// TODO a lot, including id generation, password hashing
	s.store.CreateNotebook(createRequest.Name, createRequest.ProtectionLevel, createRequest.Password)

	c.JSON(200, model.NotebookCreateResponse{
		ID: createRequest.Name,
		Response: util.Response{
			Status:  "success",
			Message: util.NotebookCreated,
		},
	})
}

func (s *Server) deleteNotebookEndpoint(c *gin.Context) {
	id := c.Param("id")
	_, exists := s.store.GetNotebook(id)

	if !exists {
		c.JSON(404, util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		})
		return
	}

	// TODO check privileges

	s.store.DeleteNotebook(id)
	c.String(200, "Ok")
}
