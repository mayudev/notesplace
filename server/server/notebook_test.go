package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/server"
	"github.com/mayudev/notesplace/server/test"
	"github.com/mayudev/notesplace/server/util"
	"github.com/stretchr/testify/assert"
)

func TestNotebookGet(t *testing.T) {
	store := &StubServerStore{
		notebooks: map[string]model.Notebook{
			"1": {
				ID:              "1",
				Name:            "Test Notebook",
				ProtectionLevel: 0,
				CreatedAt:       time.UnixMicro(0),
				UpdatedAt:       time.UnixMicro(0),
			},
		},
	}

	server := server.NewServer(store)

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

	t.Run("returns 404 with an error message on nonexistent notebook", func(t *testing.T) {

		req := test.GetAPIRequest(t, "/api/notebook/nope")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := test.DecodeJson[util.Response](t, res)
		want := util.Response{
			Status:  "error",
			Message: util.NotebookNotFound,
		}

		assert.Equal(t, 404, res.Code)
		test.AssertDeepEqual(t, got, want)
	})
}

func TestCreateNotebook(t *testing.T) {
	store := NewStubServerStore()
	server := server.NewServer(store)

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
			ID: "new notebook",
			Response: util.Response{
				Status:  "success",
				Message: util.NotebookCreated,
			},
		}

		assert.Equal(t, 200, res.Code)
		test.AssertDeepEqual(t, got, want)
	})
}
