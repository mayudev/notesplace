package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/model"
)

type Server struct {
	http.Handler
	store Store
}

type Store interface {
	GetNotebook(id string) (model.Notebook, bool)
	CreateNotebook(id string, protection uint8, hash string) error
	DeleteNotebook(id string) error
}

func NewServer(store Store) *Server {
	s := &Server{store: store}
	s.Handler = s.setupRouter()

	return s
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.Default()

	v := r.Group("/api/notebook")
	{
		v.GET("/:id", s.getNotebookEndpoint)
		v.DELETE("/:id", s.deleteNotebookEndpoint)
		v.POST("", s.createNotebookEndpoint)
	}

	return r
}
