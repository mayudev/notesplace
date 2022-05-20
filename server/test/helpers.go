package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/server"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func GetAPIRequest(t testing.TB, path string) *http.Request {
	t.Helper()

	req, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	return req
}

func PostAPIRequest(t testing.TB, path string, body string, headers http.Header) *http.Request {
	return requestWithBody(t, http.MethodPost, path, body, headers)
}

func PutAPIRequest(t testing.TB, path string, body string, headers http.Header) *http.Request {
	return requestWithBody(t, http.MethodPut, path, body, headers)
}

func requestWithBody(t testing.TB, method string, path string, body string, headers http.Header) *http.Request {
	t.Helper()

	req, err := http.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header = headers

	assert.NoError(t, err)

	return req
}

func DeleteAPIRequest(t testing.TB, path string) *http.Request {
	t.Helper()

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	assert.NoError(t, err)

	return req
}

func DecodeJson[T any](t testing.TB, res *httptest.ResponseRecorder) T {
	t.Helper()
	var got T

	decoder := jsoniter.NewDecoder(res.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&got)
	assert.NoError(t, err)

	return got
}

func EncodeJson[T any](t testing.TB, got T) string {
	t.Helper()

	encoded, err := jsoniter.MarshalToString(got)
	assert.NoError(t, err)

	return encoded
}

func AssertDeepEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid response, got %#v wanted %#v", got, want)
	}
}

func HashWithDefault(password string) string {
	hasher := auth.Hasher{Cost: bcrypt.MinCost}
	hashed, _ := hasher.HashPassword(password)
	return hashed
}

func ValidateWith(notebook string, token []byte, issuerKey string) bool {
	issuer := auth.NewIssuer(issuerKey)
	valid := issuer.ValidateNotebook(string(token), notebook)

	return valid
}

func AuthorizeFor(t testing.TB, server *server.Server, notebook, password string) string {
	t.Helper()

	authReq := GetAPIRequest(t, "/api/auth")
	authReq.Header.Add("Notebook", "protected")
	authReq.Header.Add("Password", password)
	authRes := httptest.NewRecorder()

	server.ServeHTTP(authRes, authReq)

	body, err := ioutil.ReadAll(authRes.Body)
	assert.NoError(t, err)
	assert.Equal(t, 200, authRes.Code)

	return string(body)
}
