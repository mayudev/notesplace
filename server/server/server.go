package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/repository"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	http.Handler
	store  *repository.Repository
	issuer *auth.Issuer
	hasher *auth.Hasher
	Run    func(addr ...string) (err error)
}

type ServerOptions struct {
	PrivateKey  string
	HashingCost int
}

func NewServer(store repository.Store, options ServerOptions) *Server {
	s := &Server{store: repository.NewRepository(store)}

	router := s.setupRouter()
	s.Handler = router
	s.Run = router.Run

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
