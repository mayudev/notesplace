package auth_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("hashes a password and compares it", func(t *testing.T) {
		password := "supersecret"

		hashed, err := auth.HashPassword(password)
		assert.NoError(t, err)

		success := auth.ComparePassword(hashed, password)
		assert.True(t, success)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("fails when passwords do not match", func(t *testing.T) {
		hashed, err := auth.HashPassword("qwerty")
		assert.NoError(t, err)

		success := auth.ComparePassword(hashed, "azerty")
		assert.False(t, success)
	})
}
