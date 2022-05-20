package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/server"
	"github.com/mayudev/notesplace/server/test"
	"github.com/mayudev/notesplace/server/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNotebookGet(t *testing.T) {
	password := "qwerty"

	store := &test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"1": {
				ID:              "1",
				Name:            "Test Notebook",
				ProtectionLevel: 0,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
			"protected": {
				ID:              "protected",
				Name:            "Read-protected Notebook",
				Password:        test.HashWithDefault(password),
				ProtectionLevel: 2,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
	}

	key := "unsafe_key"

	server := server.NewServer(store, server.ServerOptions{
		PrivateKey:  key,
		HashingCost: bcrypt.MinCost,
	})

	t.Run("returns information about a notebook with no protection", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/notebook/1")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := model.Notebook{
			ID:              "1",
			Name:            "Test Notebook",
			ProtectionLevel: 0,
			CreatedAt:       time.UnixMicro(0),
			UpdatedAt:       time.UnixMicro(0),
		}

		got := test.DecodeJson[model.Notebook](t, res)

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})

	t.Run("does not return information about a protected notebook to an unauthorized user", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/notebook/protected")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})

	t.Run("returns information about a protected notebook to an authorized user", func(t *testing.T) {
		token := test.AuthorizeFor(t, server, "protected", password)

		req := test.GetAPIRequest(t, "/api/notebook/protected")
		req.Header.Add("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := model.Notebook{
			ID:              "protected",
			Name:            "Read-protected Notebook",
			ProtectionLevel: 2,
			CreatedAt:       time.UnixMicro(0),
			UpdatedAt:       time.UnixMicro(0),
		}

		got := test.DecodeJson[model.Notebook](t, res)

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})

	t.Run("does not return information about a protected notebook to a user using an incorrect key", func(t *testing.T) {
		req := test.GetAPIRequest(t, "/api/notebook/protected")
		req.Header.Add("Authorization", "Bearer incorrect_token")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})

	t.Run("returns 404 with an error message on nonexistent notebook", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/notebook/nope")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[util.Response](t, res)
		want := util.Response{
			Status:  "error",
			Message: util.NotFound,
		}

		assert.Equal(t, 404, res.Code)
		test.AssertDeepEqual(t, got, want)
	})
}

func TestCreateNotebook(t *testing.T) {
	store := test.NewStubServerStore()
	server := server.NewServer(store, server.ServerOptions{
		PrivateKey:  "",
		HashingCost: bcrypt.MinCost,
	})

	t.Run("creates a notebook", func(t *testing.T) {
		body := test.EncodeJson(t, model.NotebookCreate{
			Name:            "new notebook",
			ProtectionLevel: 0,
		})

		req := test.PostAPIRequest(t, "/api/notebook", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[model.NotebookCreateResponse](t, res)

		want := model.NotebookCreateResponse{
			ID: got.ID, // ID was randomly generated
			Response: util.Response{
				Status:  "success",
				Message: util.NotebookCreated,
			},
		}

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})

	t.Run("does not allow to create a notebook with incorrect protection level", func(t *testing.T) {
		body := test.EncodeJson(t, model.NotebookCreate{
			Name:            "new notebook",
			ProtectionLevel: 3, // Invalid protection level
		})

		req := test.PostAPIRequest(t, "/api/notebook", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[util.Response](t, res)
		want := util.Response{
			Status:  "error",
			Message: util.RequestInvalid,
		}

		assert.Equal(t, 400, res.Code)
		test.AssertDeepEqual(t, got, want)
	})

	t.Run("creates a new write protected notebook", func(t *testing.T) {
		password := "supersecret"

		body := test.EncodeJson(t, model.NotebookCreate{
			Name:            "new notebook",
			ProtectionLevel: 1,
			Password:        password,
		})

		req := test.PostAPIRequest(t, "/api/notebook", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[model.NotebookCreateResponse](t, res)
		want := model.NotebookCreateResponse{
			ID: got.ID,
			Response: util.Response{
				Status:  "success",
				Message: util.NotebookCreated,
			},
		}

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)

		hashedPassword := store.Notebooks[got.ID].Password
		matches := auth.ComparePassword(hashedPassword, password)
		assert.True(t, matches)
	})
}

func TestDeleteNotebook(t *testing.T) {
	password := "unsafe_password"

	store := &test.StubServerStore{
		Notebooks: map[string]model.Notebook{
			"1": {
				ID:              "1",
				Name:            "Test Notebook",
				ProtectionLevel: 0,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
			"protected": {
				ID:              "protected",
				Name:            "Read-only notebook",
				Password:        test.HashWithDefault(password),
				ProtectionLevel: 1,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
	}

	server := server.NewServer(store, server.ServerOptions{
		PrivateKey:  "",
		HashingCost: bcrypt.MinCost,
	})

	t.Run("deletes an unprotected notebook", func(t *testing.T) {
		req := test.DeleteAPIRequest(t, "/api/notebook/1")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.NotContains(t, store.Notebooks, "1")
	})

	t.Run("refuses to delete a notebook with read-only unprivileged access", func(t *testing.T) {
		req := test.DeleteAPIRequest(t, "/api/notebook/protected")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Contains(t, store.Notebooks, "protected")
		assert.Equal(t, 401, res.Code)
	})

	t.Run("deletes a protected notebook as an authenticated user", func(t *testing.T) {
		token := test.AuthorizeFor(t, server, "protected", password)

		req := test.DeleteAPIRequest(t, "/api/notebook/protected")
		req.Header.Add("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 200, res.Code)
		assert.NotContains(t, store.Notebooks, "protected")
	})
}
