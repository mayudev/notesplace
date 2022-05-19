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
	// Notebook
	GetNotebook(id string) (model.Notebook, bool)
	CreateNotebook(id string, protection uint8, hash string) error
	DeleteNotebook(id string) error

	// Note
	GetNote(id string) (model.Note, bool)
	CreateNote(data *model.Note) error
	UpdateNote(data model.Note) (model.Note, error)
	DeleteNote(id string) error
}

func NewServer(store Store) *Server {
	s := &Server{store: store}
	s.Handler = s.setupRouter()

	return s
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.Default()

	notebook := r.Group("/api/notebook")
	{
		notebook.GET("/:id", s.getNotebookEndpoint)
		notebook.DELETE("/:id", s.deleteNotebookEndpoint)
		notebook.POST("", s.createNotebookEndpoint)
	}

	note := r.Group("/api/note")
	{
		note.GET("/:id", s.getNoteEndpoint)
		note.DELETE("/:id", s.deleteNoteEndpoint)
		note.PUT("", s.putNoteEndpoint)
	}

	return r
}
