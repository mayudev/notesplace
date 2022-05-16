package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
)

type Server struct {
	http.Handler
	store Store
}

type Store interface {
	GetNotebook(id string) (model.Notebook, bool)
	CreateNotebook(id string, protection uint8, hash string) error
}

func NewServer(store Store) *Server {
	s := &Server{store: store}
	s.Handler = s.setupRouter()

	return s
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/api/notebook/:id", func(c *gin.Context) {
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
	})

	r.POST("/api/notebook", func(c *gin.Context) {
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
	})

	return r
}
