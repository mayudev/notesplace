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
)

func TestNoteGet(t *testing.T) {
	store := &StubServerStore{
		notebooks: map[string]model.Notebook{
			"test_notebook": {
				ID:              "test_notebook",
				Name:            "Test Notebook",
				ProtectionLevel: 0,
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
		},
	}

	server := server.NewServer(store)

	t.Run("returns information about a note in an unprotected notebook", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/note/test_note")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		want := store.notes["test_note"]

		got := test.DecodeJson[model.Note](t, res)

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})
}

func TestNotePut(t *testing.T) {
	store := &StubServerStore{
		notebooks: map[string]model.Notebook{
			"test_notebook": {
				ID:              "test_notebook",
				Name:            "Test Notebook",
				ProtectionLevel: 0,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
		notes: map[string]model.Note{
			"test_note": {
				ID:         "test_note",
				NotebookID: "test_notebook",
				Title:      "Before",
				Order:      nil,
				Content:    "Before",
				CreatedAt:  time.UnixMicro(0),
				UpdatedAt:  time.UnixMicro(0),
			},
		},
	}

	server := server.NewServer(store)

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
}

func TestNoteDelete(t *testing.T) {

}
