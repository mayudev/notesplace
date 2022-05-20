package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	http.Handler
	store  Store
	issuer *auth.Issuer
	hasher *auth.Hasher
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
	PrivateKey  string
	HashingCost int
}

func NewServer(store Store, options ServerOptions) *Server {
	s := &Server{store: store}
	s.Handler = s.setupRouter()
	s.issuer = auth.NewIssuer(options.PrivateKey)

	hashingCost := bcrypt.DefaultCost

	if options.HashingCost >= bcrypt.MinCost && options.HashingCost <= bcrypt.MaxCost {
		hashingCost = options.HashingCost
	}

	s.hasher = &auth.Hasher{Cost: hashingCost}

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
