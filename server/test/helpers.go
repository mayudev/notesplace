package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
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
