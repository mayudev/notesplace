package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
)

type Server struct {
	http.Handler
	store  Store
	issuer *auth.Issuer
}

type Store interface {
	// Notebook
	GetNotebook(id string) (model.Notebook, bool)
	CreateNotebook(id string, protection auth.ProtectionLevel, hash string) error
	DeleteNotebook(id string) error

	// Note
	GetNote(id string) (model.Note, bool)
	CreateNote(data *model.Note) error
	UpdateNote(data model.Note) (model.Note, error)
	DeleteNote(id string) error
}

type ServerOptions struct {
	PrivateKey string
}

func NewServer(store Store, options ServerOptions) *Server {
	s := &Server{store: store}
	s.Handler = s.setupRouter()
	s.issuer = auth.NewIssuer(options.PrivateKey)

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

	r.GET("/api/auth", s.authenticateEndpoint)

	return r
}
