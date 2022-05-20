package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/server"
	"github.com/mayudev/notesplace/server/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var password = "unsafe_password"

var store StubServerStore = StubServerStore{
	notebooks: map[string]model.Notebook{
		"test_notebook": {
			ID:              "test_notebook",
			Name:            "Test Notebook",
			ProtectionLevel: 0,
			CreatedAt:       time.UnixMicro(0),
			UpdatedAt:       time.UnixMicro(0),
		},
		"readonly_notebook": {
			ID:              "readonly_notebook",
			Name:            "Read-only Notebook",
			ProtectionLevel: 1,
			Password:        test.HashWithDefault(password),
			CreatedAt:       time.UnixMicro(0),
			UpdatedAt:       time.UnixMicro(0),
		},
		"protected_notebook": {
			ID:              "protected_notebook",
			Name:            "Protected Notebook",
			ProtectionLevel: 2,
			Password:        test.HashWithDefault(password),
			CreatedAt:       time.UnixMicro(0),
			UpdatedAt:       time.UnixMicro(0),
		},
	},
	notes: map[string]model.Note{
		"test_note": {
			ID:         "test_note",
			NotebookID: "test_notebook",
			Title:      "Test note title",
			Order:      nil,
			Content:    "Test note contents",
			CreatedAt:  time.UnixMicro(0),
			UpdatedAt:  time.UnixMicro(0),
		},
		"readonly_note": {
			ID:         "readonly_note",
			NotebookID: "readonly_notebook",
			Title:      "Test note title",
			Order:      nil,
			Content:    "Test note contents",
			CreatedAt:  time.UnixMicro(0),
			UpdatedAt:  time.UnixMicro(0),
		},
		"protected_note": {
			ID:         "protected_note",
			NotebookID: "protected_notebook",
			Title:      "Test note title",
			Order:      nil,
			Content:    "Test note contents",
			CreatedAt:  time.UnixMicro(0),
			UpdatedAt:  time.UnixMicro(0),
		},
	},
}

func TestNoteGet(t *testing.T) {
	key := "qwerty"

	server := server.NewServer(&store, server.ServerOptions{
		PrivateKey:  key,
		HashingCost: bcrypt.MinCost,
	})

	t.Run("returns information about a note in an unprotected notebook", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/note/test_note")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := store.notes["test_note"]

		got := test.DecodeJson[model.Note](t, res)

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)

	})

	t.Run("refuses to return information about a note in a protected notebook", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/note/protected_note")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})

	t.Run("returns information about a note in a protected notebook to an authenticated user", func(t *testing.T) {
		token := test.AuthorizeFor(t, server, "protected_notebook", password)

		req := test.GetAPIRequest(t, "/api/note/protected_note")
		req.Header.Add("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := store.notes["protected_note"]

		got := test.DecodeJson[model.Note](t, res)

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})
}

func TestNotePut(t *testing.T) {
	//var store StubServerStore = store
	server := server.NewServer(&store, server.ServerOptions{
		PrivateKey: "",
	})

	t.Run("creates a note if it doesn't exist", func(t *testing.T) {
		body := test.EncodeJson(t, model.Note{
			NotebookID: "test_notebook",
			Title:      "New note",
			Order:      nil,
			Content:    "New note contents",
		})

		req := test.PutAPIRequest(t, "/api/note", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[model.Note](t, res)
		want := model.Note{
			ID:         got.ID, // ID will be generated in the backend
			NotebookID: "test_notebook",
			Title:      "New note",
			Order:      nil,
			Content:    "New note contents",
			CreatedAt:  got.CreatedAt,
			UpdatedAt:  got.UpdatedAt,
		}

		assert.Equal(t, 201, res.Code)
		test.AssertDeepEqual(t, got, want)
		test.AssertDeepEqual(t, store.notes[got.ID], want)
	})

	t.Run("updates a note's title and content in an unprotected notebook", func(t *testing.T) {
		body := test.EncodeJson(t, model.Note{
			ID:         "test_note",
			NotebookID: "test_notebook",
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
		})

		req := test.PutAPIRequest(t, "/api/note", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[model.Note](t, res)
		want := model.Note{
			ID:         "test_note",
			NotebookID: "test_notebook",
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
			CreatedAt:  time.UnixMicro(0),
			UpdatedAt:  time.UnixMicro(0),
		}

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
		test.AssertDeepEqual(t, store.notes["test_note"], want)
	})

	t.Run("updates a note's title and content in a readonly notebook to an authenticated user", func(t *testing.T) {
		token := test.AuthorizeFor(t, server, "readonly_notebook", password)

		body := test.EncodeJson(t, model.Note{
			ID:         "readonly_note",
			NotebookID: "readonly_notebook",
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
		})

		req := test.PutAPIRequest(t, "/api/note", body, http.Header{})
		req.Header.Add("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[model.Note](t, res)
		want := model.Note{
			ID:         "readonly_note",
			NotebookID: "readonly_notebook",
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
			CreatedAt:  got.CreatedAt,
			UpdatedAt:  got.UpdatedAt,
		}

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
		test.AssertDeepEqual(t, store.notes["readonly_note"], want)
	})

	t.Run("refuses an unprivileged user to update a note in a read-only notebook", func(t *testing.T) {
		body := test.EncodeJson(t, model.Note{
			ID:         "readonly_note",
			NotebookID: "readonly_notebook",
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
		})

		req := test.PutAPIRequest(t, "/api/note", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
	})

	t.Run("refuses an uprivileged user to update a note from a different notebook", func(t *testing.T) {
		body := test.EncodeJson(t, model.Note{
			ID:         "protected_note",
			NotebookID: "test_notebook", // here
			Title:      "New title",
			Order:      nil,
			Content:    "New contents",
		})

		req := test.PutAPIRequest(t, "/api/note", body, http.Header{})
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 403, res.Code)
	})
}

func TestNoteDelete(t *testing.T) {
	server := server.NewServer(&store, server.ServerOptions{
		PrivateKey: "",
	})

	t.Run("deletes a note in an unprotected notebook", func(t *testing.T) {
		req := test.DeleteAPIRequest(t, "/api/note/test_note")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.NotContains(t, store.notes, "test_note")
	})

	t.Run("refuses to delete a note in a read-only notebook", func(t *testing.T) {
		req := test.DeleteAPIRequest(t, "/api/note/readonly_note")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, 401, res.Code)
		assert.Contains(t, store.notes, "readonly_note")
	})
}
