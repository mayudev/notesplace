package server_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/server"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	password := "unsafe_password"
	hashed, _ := auth.HashPassword(password)

	store := StubServerStore{
		notebooks: map[string]model.Notebook{
			"protected": {
				ID:              "protected",
				Password:        hashed,
				ProtectionLevel: 2,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
	}

	server := server.NewServer(&store, server.ServerOptions{
		PrivateKey: "unsafe_key",
	})

	t.Run("authenticates with valid credentials", func(t *testing.T) {
		req := test.GetAPIRequest(t, "/api/auth")
		req.Header.Add("Notebook", "protected")
		req.Header.Add("Password", password)

		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 200, res.Code)
	})

	t.Run("refuses to authenticate with invalid credentials", func(t *testing.T) {
		req := test.GetAPIRequest(t, "/api/auth")
		req.Header.Add("Notebook", "protected")
		req.Header.Add("Password", "incorrect_password")

		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})
}
