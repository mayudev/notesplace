package server_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/server"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	notebook_id := "protected"
	password := "unsafe_password"

	store := test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"protected": {
				ID:              notebook_id,
				Password:        test.HashWithDefault(password),
				ProtectionLevel: 2,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
	}

	issuerKey := "unsafe_key"
	server := server.NewServer(&store, server.ServerOptions{
		PrivateKey: issuerKey,
	})

	t.Run("returns a valid JWT with valid credentials", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/auth")
		req.Header.Add("Notebook", notebook_id)
		req.Header.Add("Password", password)

		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)

		valid := test.ValidateWith(notebook_id, body, issuerKey)

		assert.True(t, valid)
		assert.Equal(t, 200, res.Code)
	})

	t.Run("refuses to authenticate with invalid credentials", func(t *testing.T) {
		req := test.GetAPIRequest(t, "/api/auth")
		req.Header.Add("Notebook", notebook_id)
		req.Header.Add("Password", "incorrect_password")

		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})
}
